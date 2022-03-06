from judgebot import private

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
WHERE chat_users.chat_id = $1
GROUP BY judge_phrases.phrase
"""

chat_user_id_query_template = "SELECT id FROM %s.%s WHERE user_id = $1 and chat_id = $2"
phrase_id_query_template = "SELECT id FROM %s.%s WHERE phrase = $1"
vote_insert_query_template = """
INSERT INTO %s.%s (vote, chat_user_id, judge_phrase_id)
VALUES ($1, $2, $3)
ON CONFLICT ON CONSTRAINT vote_pkey
DO UPDATE SET vote = $1"""
chat_user_insert_query_template = "INSERT INTO %s.%s (user_id, chat_id) VALUES ($1, $2) RETURNING id"
phrase_insert_query_template = "INSERT INTO %s.%s (phrase) VALUES ($1) RETURNING id"


def get_judge_list_query() -> str:
	return judge_list_query_template


def get_phrase_insert_query() -> str:
	return phrase_insert_query_template.format(private.DATABASE_NAME, judge_phrases)


def get_vote_insert_query() -> str:
	return vote_insert_query_template.format(private.DATABASE_NAME, votes)


def get_chat_user_id_query() -> str:
	return chat_user_id_query_template.format(private.DATABASE_NAME, chat_users)


def get_phrase_id_query() -> str:
	return phrase_id_query_template.format(private.DATABASE_NAME, judge_phrases)


def get_user_insert_query() -> str:
	return chat_user_insert_query_template.format(private.DATABASE_NAME, chat_users)
