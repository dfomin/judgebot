package sort

import (
	"sort"
	"judgebot/database"
)

type JudgePhrases []*database.JudgePhraseInfo

type lessFunc func(p1, p2 *database.JudgePhraseInfo, chatMembersCount int) bool

type sorter struct {
	judgePhrases     []database.JudgePhraseInfo
	chatMembersCount int
	less             lessFunc
}

func (s *sorter) Len() int {
	return len(s.judgePhrases)
}

func (s *sorter) Swap(i, j int) {
	s.judgePhrases[i], s.judgePhrases[j] = s.judgePhrases[j], s.judgePhrases[i]
}

func (s *sorter) Less(i, j int) bool {
	p, q := &s.judgePhrases[i], &s.judgePhrases[j]
	return s.less(p, q, s.chatMembersCount)
}

func (s *sorter) Sort(judgePhrases []database.JudgePhraseInfo) {
	s.judgePhrases = judgePhrases
	sort.Sort(s)
}

func OrderBy(less lessFunc, chatMembersCount int) *sorter {
	return &sorter{
		less: less,
		chatMembersCount: chatMembersCount,
	}
}
