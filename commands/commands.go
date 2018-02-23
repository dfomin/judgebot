package commands

import (
	"judgebot/database"
)

var _dbc *database.Controller = nil

func dbc() *database.Controller {
	if _dbc == nil {
		_dbc = database.Init()
	}

	return _dbc
}

func Judge(names []string) string {
	result := ""
	for _, name := range names {
		result += name + "\n"
	}

	return result
}

func JudgeList() []string {
	return dbc().JudgeList()
}
