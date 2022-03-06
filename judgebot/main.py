from telegram import Update
from telegram.ext import Updater, Dispatcher, CallbackContext, CommandHandler

from judgebot.private import TOKEN


def start(update: Update, context: CallbackContext):
    context.bot.send_message(chat_id=update.effective_chat.id, text="👋🏻")


def main():
    updater = Updater(TOKEN, use_context=True)

    updater.start_webhook(listen="127.0.0.1",
                          port=5002,
                          url_path=f"{TOKEN}",
                          webhook_url=f"https://dfomin.com:443/{TOKEN}")
    dispatcher: Dispatcher = updater.dispatcher
    dispatcher.add_handler(CommandHandler("start", start))


if __name__ == "__main__":
    main()