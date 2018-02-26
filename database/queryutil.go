package database

import (
	"fmt"
	"judgebot/private"
)

const users = `users`
const judgePhrases = `judge_phrases`
const votes = `votes`

const judgeListQueryTemplate = `SELECT phrase FROM %s.%s`
const userIDQueryTemplate = `SELECT id FROM %s.%s WHERE telegram_id = $1`
const phraseIDQueryTemplate = `SELECT id FROM %s.%s WHERE phrase = $1`
const voteInsertQueryTemplate = `INSERT INTO %s.%s (vote, user_id, judge_phrase_id) VALUES ($1, $2, $3) 
ON CONFLICT ON CONSTRAINT vote_pkey DO UPDATE SET vote = $1`
const userInsertQueryTemplate = `INSERT INTO %s.%s (telegram_id) VALUES ($1) RETURNING id`
const phraseInsertQueryTemplate = `INSERT INTO %s.%s (phrase) VALUES ($1) RETURNING id`

func getJudgeListQuery() string {
	return fmt.Sprintf(judgeListQueryTemplate, private.DatabaseName, judgePhrases)
}

func getPhraseInsertQuery() string {
	return fmt.Sprintf(phraseInsertQueryTemplate, private.DatabaseName, judgePhrases)
}

func getVoteInsertQuery() string {
	return fmt.Sprintf(voteInsertQueryTemplate, private.DatabaseName, votes)
}

func getUserIDQuery() string {
	return fmt.Sprintf(userIDQueryTemplate, private.DatabaseName, users)
}

func getPhraseIDQuery() string {
	return fmt.Sprintf(phraseIDQueryTemplate, private.DatabaseName, judgePhrases)
}

func getUserInsertQuery() string {
	return fmt.Sprintf(userInsertQueryTemplate, private.DatabaseName, users)
}
