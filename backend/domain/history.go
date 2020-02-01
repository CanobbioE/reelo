package domain

// History is a map of results details indexed by year
type History map[int]struct {
	Partecipation Partecipation `json:"partecipation"`
	MaxExercises  int           `json:"eMax"`
	MaxScore      int           `json:"dMax"`
}

// playerHistoryByPlayerNameAndSurname(ctx context.Context, n,s string) (History, error)
// analysisHistoryByPlayerID(ctx context.Context, id int) (History, []int, error)
// analysisHistoryByPlayerNameAndSurname(ctx context.Context, n, s string) (History, []int, error)

// HistorySwitcheroo(ctx context.Context, oldID, newID int, newHistory []History) error
