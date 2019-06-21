package dto

// Rank represent one entry of a ranking
type Rank struct {
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Category string  `json:"category"`
	Reelo    float64 `json:"reelo"`
}
