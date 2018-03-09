package database

import (
	"fmt"
	"judgebot/private"
)

const chatUsers = `chat_users`
const judgePhrases = `judge_phrases`
const votes = `votes`

const judgeListQueryTemplate = `
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
`

const chatUserIDQueryTemplate = `SELECT id FROM %s.%s WHERE user_id = $1 and chat_id = $2`
const phraseIDQueryTemplate = `SELECT id FROM %s.%s WHERE phrase = $1`
const voteInsertQueryTemplate = `INSERT INTO %s.%s (vote, chat_user_id, judge_phrase_id) VALUES ($1, $2, $3)
ON CONFLICT ON CONSTRAINT vote_pkey DO UPDATE SET vote = $1`
const chatUserInsertQueryTemplate = `INSERT INTO %s.%s (user_id, chat_id) VALUES ($1, $2) RETURNING id`
const phraseInsertQueryTemplate = `INSERT INTO %s.%s (phrase) VALUES ($1) RETURNING id`

func getJudgeListQuery() string {
	return judgeListQueryTemplate
}

func getPhraseInsertQuery() string {
	return fmt.Sprintf(phraseInsertQueryTemplate, private.DatabaseName, judgePhrases)
}

func getVoteInsertQuery() string {
	return fmt.Sprintf(voteInsertQueryTemplate, private.DatabaseName, votes)
}

func getChatUserIDQuery() string {
	return fmt.Sprintf(chatUserIDQueryTemplate, private.DatabaseName, chatUsers)
}

func getPhraseIDQuery() string {
	return fmt.Sprintf(phraseIDQueryTemplate, private.DatabaseName, judgePhrases)
}

func getUserInsertQuery() string {
	return fmt.Sprintf(chatUserInsertQueryTemplate, private.DatabaseName, chatUsers)
}
