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

func Init() *Controller {
	info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=verify-full", private.DatabaseUser, private.DatabasePassword, private.DatabaseName)

	db, err := sql.Open("postgres", info)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &Controller{DataBase: db}
}

func (dbc *Controller) JudgeList() []string {
	query := "SELECT phrase FROM " + private.DatabaseName + ".judge_phrases"
	rows, err := dbc.DataBase.Query(query)
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
