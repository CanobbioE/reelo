package db

// TODO we could implement something like a model mapper
// using golang tags for easier conversion

// Result represents the result entity in the database
// TODO Since all the following functions recovers information that could be
// obtained by manipulating the Result array returned by GetResults,
// we should refactor all of this as a datatype. Be careul to not overdo it:
// using SQL is really efficient compared to iterating over data structures
// avg, min and max functions should be DB based.
type Result struct {
	Time      int
	Exercises int
	Score     int
	Year      int
	Category  string
}

// Player represents the player entity as it is stored in the db
type Player struct {
	Name    string
	Surname string
	Reelo   int
}

// Costants represents all the costants used to calculate the reelo score
type Costants struct {
	StartingYear           int     `json:"year"`
	ExercisesCostant       float64 `json:"ex"`
	PFinal                 float64 `json:"final"`
	MultiplicativeFactor   float64 `json:"mult"`
	AntiExploit            float64 `json:"exp"`
	NoPartecipationPenalty float64 `json:"np"`
}

// GameInfo contains the data about a certain game
type GameInfo struct {
	Year     int
	Category string
	Start    int
	End      int
}
