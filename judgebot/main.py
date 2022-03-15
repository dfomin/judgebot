from typing import List

import psycopg
from telegram import Update
from telegram.ext import Updater, Dispatcher, CallbackContext, CommandHandler

from judgebot.private import TOKEN, DATABASE_NAME, DATABASE_PASSWORD, DATABASE_USER
from judgebot.sql.queries import get_judge_list_query


def start(update: Update, context: CallbackContext):
    context.bot.send_message(chat_id=update.effective_chat.id, text="👋🏻")


def help(update: Update, context: CallbackContext):
    context.bot.send_message(chat_id=update.effective_chat.id, text="😐")


def judge(update: Update, context: CallbackContext):
    # context.bot.send_message(chat_id=update.effective_chat.id, text="😐")
    phrases = applicable_judge_list(update.effective_chat.id, 0)
    context.bot.send_message(chat_id=update.effective_chat.id, text=phrases[0])


def applicable_judge_list(chat_id: int, chat_members_count: int) -> List[str]:
    phrases = get_sorted_judge_phrases(chat_id, chat_members_count)
    return phrases


def get_sorted_judge_phrases(chat_id: int, chat_members_count: int) -> List[str]:
    with psycopg.connect(f"dbname={DATABASE_NAME} password={DATABASE_PASSWORD} user={DATABASE_USER} sslmode=disable") as conn:
        with conn.cursor() as cur:
            cur.execute(get_judge_list_query(), f"{chat_id}")
            return cur.fetchall()


def main():
    updater = Updater(TOKEN, use_context=True)

    updater.start_webhook(listen="127.0.0.1",
                          port=5002,
                          url_path=TOKEN,
                          webhook_url=f"https://dfomin.com:443/{TOKEN}")
    dispatcher: Dispatcher = updater.dispatcher
    dispatcher.add_handler(CommandHandler("start", start))
    dispatcher.add_handler(CommandHandler("help", help))
    dispatcher.add_handler(CommandHandler("judge", judge))


if __name__ == "__main__":
    main()
