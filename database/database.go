package database

import (
	"database/sql"
	"fmt"
	"judgebot/private"
	"log"

	_ "github.com/lib/pq"
)

type Controller struct {
	DataBase *sql.DB
}

type JudgePhraseInfo struct {
	Phrase   string
	Voteup   int
	Votedown int
}

func Init() *Controller {
	//info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=verify-full", private.DatabaseUser, private.DatabasePassword, private.DatabaseName)
	info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", private.DatabaseUser, private.DatabasePassword, private.DatabaseName)

	database, err := sql.Open("postgres", info)
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
	err := dbc.DataBase.QueryRow(getUserIDQuery(), telegramID).Scan(&userID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		err = dbc.DataBase.QueryRow(getUserInsertQuery(), telegramID).Scan(&userID)
		if err != nil {
			log.Fatal(err)
		}
	}

	return userID
}

func (dbc *Controller) getPhraseID(phrase string) int {
	var phraseID int
	err := dbc.DataBase.QueryRow(getPhraseIDQuery(), phrase).Scan(&phraseID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		err = dbc.DataBase.QueryRow(getPhraseInsertQuery(), phrase).Scan(&phraseID)
		if err != nil {
			log.Fatal(err)
		}
	}

	return phraseID
}

func (dbc *Controller) JudgeVote(telegramID int, phrase string, vote bool) {
	userID := dbc.getUserID(telegramID)
	phraseID := dbc.getPhraseID(phrase)
	_, err := dbc.DataBase.Query(getVoteInsertQuery(), vote, userID, phraseID)
	if err != nil {
		log.Fatal(err)
	}
}

func (dbc *Controller) JudgeList() []JudgePhraseInfo {
	rows, err := dbc.DataBase.Query(getJudgeListQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var phrases []JudgePhraseInfo
	for rows.Next() {
		var phrase string
		var voteup int
		var votedown int
		if err := rows.Scan(&phrase, &voteup, &votedown); err != nil {
			log.Fatal(err)
		}
		phrases = append(phrases, JudgePhraseInfo{phrase, voteup, votedown})
	}

	return phrases
}
