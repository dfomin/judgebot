import random
from dataclasses import dataclass
from enum import Enum, auto
from time import time
from typing import List

import psycopg
from telegram import Update
from telegram.ext import Updater, Dispatcher, CallbackContext, CommandHandler

from judgebot.private import TOKEN, DATABASE_NAME, DATABASE_PASSWORD, DATABASE_USER
from judgebot.sql.queries import get_judge_list_query


class PhraseStatus(Enum):
    ACCEPTED = auto()
    IN_PROGRESS = auto()
    REJECTED = auto()


@dataclass
class Phrase:
    text: str
    vote_up: int
    vote_down: int
    status: PhraseStatus
    sort_value: float


def start(update: Update, context: CallbackContext):
    context.bot.send_message(chat_id=update.effective_chat.id, text="👋🏻")


def help(update: Update, context: CallbackContext):
    context.bot.send_message(chat_id=update.effective_chat.id, text="😐")


def judge(update: Update, context: CallbackContext):
    parts = update.message.text.split()
    if len(parts) <= 1:
        context.bot.send_message(chat_id=update.effective_chat.id, text="😐")
    else:
        chat_members_count = context.bot.get_chat_member_count(update.effective_chat.id) - 1
        phrases = applicable_judge_list(update.effective_chat.id, chat_members_count)
        phrases = [phrase for phrase in phrases if phrase.status == PhraseStatus.ACCEPTED]
        result = ""
        for part in parts[1:]:
            phrase = random.choice(phrases)
            if len(result) > 0:
                result += "\n"
            result += phrase.text.replace("%", part)
        context.bot.send_message(chat_id=update.effective_chat.id, text=result)


def judge_list(update: Update, context: CallbackContext):
    chat_members_count = context.bot.get_chat_member_count(update.effective_chat.id) - 1
    phrases = applicable_judge_list(update.effective_chat.id, chat_members_count)
    current_status = None
    result = ""
    for phrase in sorted(phrases, key=lambda x: x.sort_value):
        if current_status is None:
            current_status = phrase.status
        elif current_status != phrase.status:
            result += "\n"
            current_status = phrase.status
        if phrase.status == PhraseStatus.ACCEPTED:
            result += "+ "
        else:
            result += "- "
        result += f"{phrase.vote_up} {phrase.vote_down} {phrase.text}\n"
    context.bot.send_message(chat_id=update.effective_chat.id, text=result)


def applicable_judge_list(chat_id: int, chat_members_count: int) -> List[Phrase]:
    phrases = get_sorted_judge_phrases(chat_id, chat_members_count)
    return phrases


def get_sorted_judge_phrases(chat_id: int, chat_members_count: int) -> List[Phrase]:
    with psycopg.connect(f"dbname={DATABASE_NAME} password={DATABASE_PASSWORD} user={DATABASE_USER} sslmode=disable") as conn:
        with conn.cursor() as cur:
            cur.execute(get_judge_list_query(), (chat_id,))
            rows = cur.fetchall()
            phrases = []
            for row in rows:
                text = row[0]
                vote_up = row[1]
                vote_down = row[2]
                sort_value = -(vote_up - vote_down) * 100 - vote_up     # Will not work for huge amount of votes
                if vote_up + vote_down >= chat_members_count / 2:
                    if vote_up - vote_down >= chat_members_count / 3:
                        status = PhraseStatus.ACCEPTED
                        sort_value -= 1000000
                    else:
                        status = PhraseStatus.REJECTED
                else:
                    status = PhraseStatus.IN_PROGRESS
                    sort_value -= 10000
                phrases.append(Phrase(text, vote_up, vote_down, status, sort_value))
            return phrases


def main():
    random.seed(time())

    updater = Updater(TOKEN, use_context=True)

    updater.start_webhook(listen="127.0.0.1",
                          port=5002,
                          url_path=TOKEN,
                          webhook_url=f"https://dfomin.com:443/{TOKEN}")
    dispatcher: Dispatcher = updater.dispatcher
    dispatcher.add_handler(CommandHandler("start", start))
    dispatcher.add_handler(CommandHandler("help", help))
    dispatcher.add_handler(CommandHandler("judge", judge))
    dispatcher.add_handler(CommandHandler("judgelist", judge_list))


if __name__ == "__main__":
    main()
