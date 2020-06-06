package memdb

// InMemoryDB represents a temporary database that reflects part of the
// real database' schema.
// The in-memory database uses addresses as reference, instead of IDs.
type InMemoryDB struct {
	Players        []Player
	Games          []Game
	Participations []Participation
	Results        []Result
}

// Player contains the player's name and surname.
type Player struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// Game contains a game's details
type Game struct {
	Year     int    `json:"year"`
	Category string `json:"category"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
}

// Result contains a player's result values.
type Result struct {
	Exercises int `json:"exercises"`
	Time      int `json:"time"`
	Score     int `json:"score"`
	Position  int `json:"position"`
}

// Participation is the relationship between players, games and results.
type Participation struct {
	Player *Player `json:"player"`
	Game   *Game   `json:"game"`
	Result *Result `json:"result"`
	City   string  `json:"city"`
}

// New creates a new instance for the in-memory databse
func New() *InMemoryDB {
	return &InMemoryDB{
		Players:        []Player{},
		Games:          []Game{},
		Participations: []Participation{},
		Results:        []Result{},
	}
}

// ContainsPlayer returns whether or not a player is already in the db.Players array.
func (db *InMemoryDB) ContainsPlayer(name, surname string) bool {
	for _, p := range db.Players {
		if p.Name == name && p.Surname == surname {
			return true
		}
	}
	return false
}

// AddPlayer appends a new player to the players array and returns its memory address.
func (db *InMemoryDB) AddPlayer(name, surname string) *Player {
	db.Players = append(db.Players, Player{name, surname})
	return &db.Players[len(db.Players)-1]
}

// AddGame appends a new game to the games array and returns its memory address.
func (db *InMemoryDB) AddGame(year int, category string, start, end int) *Game {
	db.Games = append(db.Games, Game{year, category, start, end})
	return &db.Games[len(db.Games)-1]
}

// AddResult appends a new result to the results array and returns its memory address.
func (db *InMemoryDB) AddResult(exercises, time, points, position int) *Result {
	db.Results = append(db.Results, Result{exercises, time, points, position})
	return &db.Results[len(db.Results)-1]
}

// AddParticipation appends a new Participation to the Participations array and returns its memory address.
func (db *InMemoryDB) AddParticipation(p *Player, g *Game, r *Result, city string) *Participation {
	db.Participations = append(db.Participations, Participation{p, g, r, city})
	return &db.Participations[len(db.Participations)-1]
}
