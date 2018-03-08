package commands

import (
	"judgebot/database"
	"math/rand"
	"strings"
	"strconv"
)

const sharpSymbol = "#"

var _dbc *database.Controller = nil

func dbc() *database.Controller {
	if _dbc == nil {
		_dbc = database.Init()
	}

	return _dbc
}

func Judge(names []string, chatMembersCount int) string {
	result := ""
	judgePhrases := applicableJudgeList(chatMembersCount)
	if len(judgePhrases) == 0 {
		return "There are no judge phrases\n"
	}
	for _, name := range names {
		phrase := judgePhrases[rand.Intn(len(judgePhrases))]
		result += strings.Replace(phrase.Phrase, sharpSymbol, name, -1) + "\n"
	}
	return result
}

func JudgeList(chatMembersCount int) string {
	result := ""
	judgePhrases := dbc().JudgeList()
	for _, judgePhrase := range judgePhrases {
		prefix := "- "
		if inFavor(judgePhrase, chatMembersCount) {
			prefix = "+ "
		}
		result += prefix + judgePhrase.Phrase + " " + strconv.Itoa(judgePhrase.Voteup) + " " + strconv.Itoa(judgePhrase.Votedown) + "\n"
	}
	return result
}

func JudgeVote(userID int, phrase string, vote bool) {
	dbc().JudgeVote(userID, phrase, vote)
}

// N-chatMembersCount, x-in favor, y-against
// x-y >= N/3 && x+y >= N/2
func inFavor(judgePhrase database.JudgePhraseInfo, chatMembersCount int) bool {
	return float64(judgePhrase.Voteup-judgePhrase.Votedown) >= float64(chatMembersCount)/3 && float64(judgePhrase.Voteup+judgePhrase.Votedown) >= float64(chatMembersCount)/2
}

func applicableJudgeList(chatMembersCount int) []database.JudgePhraseInfo {
	allPhrases := dbc().JudgeList()
	var phrases []database.JudgePhraseInfo
	for _, phrase := range allPhrases {
		if inFavor(phrase, chatMembersCount) && strings.Contains(phrase.Phrase, sharpSymbol) {
			phrases = append(phrases, phrase)
		}
	}
	return phrases
}
