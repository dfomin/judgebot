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

func (dbc *Controller) getChatUserID(userID int, chatID int64) int {
	var chatUserID int
	err := dbc.DataBase.QueryRow(getChatUserIDQuery(), userID, chatID).Scan(&chatUserID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
		err = dbc.DataBase.QueryRow(getUserInsertQuery(), userID, chatID).Scan(&chatUserID)
		if err != nil {
			log.Fatal(err)
		}
	}

	return chatUserID
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

func (dbc *Controller) JudgeVote(userID int, chatID int64, phrase string, vote bool) {
	chatUserID := dbc.getChatUserID(userID, chatID)
	phraseID := dbc.getPhraseID(phrase)
	rows, err := dbc.DataBase.Query(getVoteInsertQuery(), vote, chatUserID, phraseID)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (dbc *Controller) JudgeList(chatID int64) []JudgePhraseInfo {
	rows, err := dbc.DataBase.Query(getJudgeListQuery(), chatID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var phrases []JudgePhraseInfo
	for rows.Next() {
		var phrase string
		var voteUp, voteDown int
		if err := rows.Scan(&phrase, &voteUp, &voteDown); err != nil {
			log.Fatal(err)
		}
		phrases = append(phrases, JudgePhraseInfo{phrase, voteUp, voteDown})
	}

	return phrases
}
