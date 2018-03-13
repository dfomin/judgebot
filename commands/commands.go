package commands

import (
	"judgebot/database"
	"math/rand"
	"strconv"
	"strings"
	"judgebot/sort"
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
	judgePhrases := getSortedJudgePhrases(chatID, chatMembersCount)

	others := false
	forceToVote := false
	for _, judgePhrase := range judgePhrases {
		prefix := "- "
		if inFavor(judgePhrase, chatMembersCount) {
			prefix = "+ "
		} else if !allVotesMoreThanMin(judgePhrase, chatMembersCount) {
			if len(result) > 0 && !forceToVote {
				forceToVote = true
				result += "\n"
			}
		} else if !votesForMoreThanMin(judgePhrase, chatMembersCount) {
			if len(result) > 0 && !others {
				others = true
				result += "\n"
			}
		}

		result += prefix + strconv.Itoa(judgePhrase.Voteup) + " " + strconv.Itoa(judgePhrase.Votedown) + " " + judgePhrase.Phrase + "\n"
	}
	return result
}

func JudgeVote(userID int, chatID int64, phrase string, vote bool) {
	dbc().JudgeVote(userID, chatID, phrase, vote)
}

func inFavor(judgePhrase database.JudgePhraseInfo, chatMembersCount int) bool {
	return votesForMoreThanMin(judgePhrase, chatMembersCount) && allVotesMoreThanMin(judgePhrase, chatMembersCount)
}

// N-chatMembersCount, x-in favor, y-against
// x-y >= N/3
func votesForMoreThanMin(judgePhrase database.JudgePhraseInfo, chatMembersCount int) bool {
	votesFor := float64(judgePhrase.Voteup - judgePhrase.Votedown)
	votesForMin := float64(chatMembersCount) / 3

	return votesFor >= votesForMin
}

// N-chatMembersCount, x-in favor, y-against
// x+y >= N/2
func allVotesMoreThanMin(judgePhrase database.JudgePhraseInfo, chatMembersCount int) bool {
	allVotes := float64(judgePhrase.Voteup + judgePhrase.Votedown)
	allVotesMin := float64(chatMembersCount) / 2

	return allVotes >= allVotesMin
}

func applicableJudgeList(chatID int64, chatMembersCount int) []database.JudgePhraseInfo {
	allPhrases := getSortedJudgePhrases(chatID, chatMembersCount)

	var phrases []database.JudgePhraseInfo
	for _, phrase := range allPhrases {
		if inFavor(phrase, chatMembersCount) && strings.Contains(phrase.Phrase, placeholderSymbol) {
			phrases = append(phrases, phrase)
		}
	}
	return phrases
}

func getSortedJudgePhrases(chatID int64, chatMembersCount int) []database.JudgePhraseInfo {
	phrases := dbc().JudgeList(chatID)
	sortFunc := func(j1, j2 *database.JudgePhraseInfo, chatMembersCount int) bool {
		if inFavor(*j1, chatMembersCount) && inFavor(*j2, chatMembersCount) {
			return j1.Voteup-j1.Votedown > j2.Voteup-j2.Votedown
		}
		if inFavor(*j1, chatMembersCount) && !inFavor(*j2, chatMembersCount) {
			return true
		}
		if !inFavor(*j1, chatMembersCount) && inFavor(*j2, chatMembersCount) {
			return false
		}
		if allVotesMoreThanMin(*j1, chatMembersCount) && allVotesMoreThanMin(*j2, chatMembersCount) {
			if j1.Voteup-j1.Votedown > j2.Voteup-j2.Votedown {
				return true
			}
			if j1.Voteup-j1.Votedown < j2.Voteup-j2.Votedown {
				return false
			}
			return j1.Votedown < j2.Votedown
		}
		if allVotesMoreThanMin(*j1, chatMembersCount) && !allVotesMoreThanMin(*j2, chatMembersCount) {
			return false
		}
		if !allVotesMoreThanMin(*j1, chatMembersCount) && allVotesMoreThanMin(*j2, chatMembersCount) {
			return true
		}
		if votesForMoreThanMin(*j1, chatMembersCount) && votesForMoreThanMin(*j2, chatMembersCount) {
			return j1.Voteup+j1.Votedown > j2.Voteup+j2.Votedown
		}
		if votesForMoreThanMin(*j1, chatMembersCount) && !votesForMoreThanMin(*j2, chatMembersCount) {
			return true
		}
		return false
	}

	sort.OrderBy(sortFunc, chatMembersCount).Sort(phrases)
	return phrases
}
