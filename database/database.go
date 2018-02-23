package database

import (
	"database/sql"
	"fmt"
	"judgebot/private"
	"log"

	_ "github.com/lib/pq"
)

const judgeListQuery = `SELECT phrase FROM judgebot.judge_phrases`
const judgePhraseQuery = `SELECT id FROM judgebot.judge_phrases WHERE phrase LIKE '$1'`
const judgePhraseInsertQuery = `INSERT INTO judgebot.judge_phrases (phrase) VALUES ($1)`
const voteForPhraseQuery = `INSERT INTO judgebot.votes (vote, user_id, judge_phrase_id) VALUES ($1, $2, $3)`
const userIDQuery = `SELECT id FROM judgebot.users WHERE telegram_id = $1`
const phraseIDQuery = `SELECT id FROM judgebot.judge_phrases WHERE phrase LIKE $1`

type Controller struct {
	DataBase *sql.DB
}

func InitDatabase(name string) *Controller {
	databaseInfo := fmt.Sprintf("user=judgebot password=%s dbname=judgebot sslmode=disable", private.DatabasePassword)

	database, err := sql.Open("postgres", databaseInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Controller{DataBase: database}
}

func (dbc *Controller) getUserID(telegramID int) int {
	var userID int
	err := dbc.DataBase.QueryRow(userIDQuery, telegramID).Scan(&userID)
	if err != nil {
		log.Fatal(err)
	}

	return userID
}

func (dbc *Controller) getPhraseID(phrase string) int {
	var phraseID int
	err := dbc.DataBase.QueryRow(phraseIDQuery, phrase).Scan(&phraseID)
	if err != nil {
		log.Fatal(err)
	}

	return phraseID
}

func (dbc *Controller) JudgeVote(telegramID int, phrase string, vote bool) {
	userID := dbc.getUserID(telegramID)
	phraseID := dbc.getPhraseID(phrase)
	_, err := dbc.DataBase.Query(voteForPhraseQuery, vote, userID, phraseID)
	if err != nil {
		log.Fatal(err)
	}
}

func (dbc *Controller) JudgeList() []string {
	rows, err := dbc.DataBase.Query(judgeListQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var phrases []string
	for rows.Next() {
		var phrase string
		if err := rows.Scan(&phrase); err != nil {
			log.Fatal(err)
		}
		phrases = append(phrases, phrase)
	}

	return phrases
}
