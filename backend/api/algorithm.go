package api

// Costants represent all the possible variables used when calculating the reelo
type Costants struct {
	StartingYear           int     `json:"year"`
	ExercisesCostant       float64 `json:"ex"`
	PFinal                 float64 `json:"final"`
	MultiplicativeFactor   float64 `json:"mult"`
	AntiExploit            float64 `json:"exp"`
	NoPartecipationPenalty float64 `json:"np"`
}
