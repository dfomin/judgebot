package database

import (
	"database/sql"
	"fmt"
	"judgebot/private"
	"log"
)

type Controller struct {
	DataBase *sql.DB
}

func InitDatabase(name string) *Controller {
	databaseInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=verify-full", private.DatabaseUser, private.DatabasePassword, private.DatabaseName)

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
