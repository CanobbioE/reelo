package dto

// Costants represent all the possible variables used when calculating the reelo
type Costants struct {
	StartingYear           int     `json:"year",omitempty`
	ExercisesCostant       float64 `json:"ex",omitempty`
	PFinal                 float64 `json:"final",omitempty`
	MultiplicativeFactor   float64 `json:"mult",omitempty`
	AntiExploit            float64 `json:"exp",omitempty`
	NoPartecipationPenalty float64 `json:"np",omitempty`
}
