package commands

import (
	"judgebot/database"
	"math/rand"
	"strconv"
	"strings"
)

const placeholderSymbol = "%"

var _dbc *database.Controller = nil

func dbc() *database.Controller {
	if _dbc == nil {
		_dbc = database.Init()
	}

	return _dbc
}

func Judge(names []string, chatID int64, chatMembersCount int) string {
	result := ""
	judgePhrases := applicableJudgeList(chatID, chatMembersCount)
	if len(judgePhrases) == 0 {
		return "There are no judge phrases\n"
	}
	for _, name := range names {
		phrase := judgePhrases[rand.Intn(len(judgePhrases))]
		result += strings.Replace(phrase.Phrase, placeholderSymbol, name, -1) + "\n"
	}
	return result
}

func JudgeList(chatID int64, chatMembersCount int) string {
	result := ""
	judgePhrases := dbc().JudgeList(chatID)
	for _, judgePhrase := range judgePhrases {
		prefix := "- "
		if inFavor(judgePhrase, chatMembersCount) {
			prefix = "+ "
		}
		result += prefix + judgePhrase.Phrase + " " + strconv.Itoa(judgePhrase.Voteup) + " " + strconv.Itoa(judgePhrase.Votedown) + "\n"
	}
	return result
}

func JudgeVote(userID int, chatID int64, phrase string, vote bool) {
	dbc().JudgeVote(userID, chatID, phrase, vote)
}

// N-chatMembersCount, x-in favor, y-against
// x-y >= N/3 && x+y >= N/2
func inFavor(judgePhrase database.JudgePhraseInfo, chatMembersCount int) bool {
	votesFor := float64(judgePhrase.Voteup - judgePhrase.Votedown)
	votesForMin := float64(chatMembersCount) / 3
	allVotes := float64(judgePhrase.Voteup + judgePhrase.Votedown)
	allVotesMin := float64(chatMembersCount) / 2

	return votesFor >= votesForMin && allVotes >= allVotesMin
}

func applicableJudgeList(chatID int64, chatMembersCount int) []database.JudgePhraseInfo {
	allPhrases := dbc().JudgeList(chatID)
	var phrases []database.JudgePhraseInfo
	for _, phrase := range allPhrases {
		if inFavor(phrase, chatMembersCount) && strings.Contains(phrase.Phrase, placeholderSymbol) {
			phrases = append(phrases, phrase)
		}
	}
	return phrases
}
