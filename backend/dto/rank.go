package dto

// Rank represents one entry of a ranking
type Rank struct {
	Name     string        `json:"name"`
	Surname  string        `json:"surname"`
	Category string        `json:"category"`
	Reelo    float64       `json:"reelo"`
	History  PlayerHistory `json:"history"`
}

// PlayerHistory is a map of years and histories.
type PlayerHistory map[int]History

// History represents one year's details for a player
type History struct {
	PseudoReelo  float64 `json:"pseudoReelo"`
	Category     string  `json:"category"`
	Exercises    int     `json:"e"`
	MaxExercises int     `json:"eMax"`
	Score        int     `json:"d"`
	MaxScore     int     `json:"dMax"`
	Time         int     `json:"time"`
	Position     int     `json:"position"`
}
