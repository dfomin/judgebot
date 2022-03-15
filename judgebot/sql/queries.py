from judgebot import private
from judgebot.private import DATABASE_NAME

chat_users = "chat_users"
judge_phrases = "judge_phrases"
votes = "votes"

judge_list_query_template = """
SELECT judge_phrases.phrase,
SUM(case when votes.vote = true then 1 else 0 end),
SUM(case when votes.vote = false then 1 else 0 end)
FROM judgebot.judge_phrases
INNER JOIN judgebot.votes
ON judge_phrases.id = votes.judge_phrase_id
INNER JOIN judgebot.chat_users
ON votes.chat_user_id = chat_users.id
WHERE chat_users.chat_id = %s
GROUP BY judge_phrases.phrase
"""

chat_user_id_query_template = f"SELECT id FROM {DATABASE_NAME}.{chat_users} WHERE user_id = %s and chat_id = %s"
phrase_id_query_template = f"SELECT id FROM {DATABASE_NAME}.{judge_phrases} WHERE phrase = %s"
vote_insert_query_template = f"""
INSERT INTO {DATABASE_NAME}.{votes} (vote, chat_user_id, judge_phrase_id)
VALUES (%s, %s, %s)
ON CONFLICT ON CONSTRAINT vote_pkey
DO UPDATE SET vote = %s"""
chat_user_insert_query_template = f"""
INSERT INTO {DATABASE_NAME}.{chat_users} (user_id, chat_id)
VALUES (%s, %s)
RETURNING id"""
phrase_insert_query_template = "INSERT INTO {DATABASE_NAME}.{chat_users} (phrase) VALUES (%s) RETURNING id"


def get_judge_list_query() -> str:
    return judge_list_query_template


def get_phrase_insert_query() -> str:
    return phrase_insert_query_template


def get_vote_insert_query() -> str:
    return vote_insert_query_template


def get_chat_user_id_query() -> str:
    return chat_user_id_query_template


def get_phrase_id_query() -> str:
    return phrase_id_query_template


def get_user_insert_query() -> str:
    return chat_user_insert_query_template.format(private.DATABASE_NAME, chat_users)
