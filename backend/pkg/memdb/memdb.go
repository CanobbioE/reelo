package memdb

import (
	"fmt"
)

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
	ID      *Player
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// Game contains a game's details
type Game struct {
	ID       *Game
	Year     int    `json:"year"`
	Category string `json:"category"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
}

// Result contains a player's result values.
type Result struct {
	ID        *Result
	Exercises int `json:"exercises"`
	Time      int `json:"time"`
	Score     int `json:"score"`
	Position  int `json:"position"`
}

// Participation is the relationship between players, games and results.
type Participation struct {
	ID     *Participation
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
	p := Player{nil, name, surname}
	p.ID = &p
	db.Players = append(db.Players, p)
	return &p
}

// AddGame appends a new game to the games array and returns its memory address.
func (db *InMemoryDB) AddGame(year int, category string, start, end int) *Game {
	g := Game{nil, year, category, start, end}
	g.ID = &g
	db.Games = append(db.Games, g)
	return &g
}

// AddResult appends a new result to the results array and returns its memory address.
func (db *InMemoryDB) AddResult(exercises, time, points, position int) *Result {
	r := Result{nil, exercises, time, points, position}
	r.ID = &r
	db.Results = append(db.Results, r)
	return &r
}

// AddParticipation appends a new Participation to the Participations array and returns its memory address.
func (db *InMemoryDB) AddParticipation(p *Player, g *Game, r *Result, city string) *Participation {
	part := Participation{nil, p, g, r, city}
	part.ID = &part
	db.Participations = append(db.Participations, part)
	return &part
}

// Show prints the whole database in a human readable format, used for debugging
func (db *InMemoryDB) Show() {
	fmt.Println("-------- Giocatori --------")
	for _, player := range db.Players {
		fmt.Printf("%p\t\t| %v\t\t| %v\n", player.ID, player.Name, player.Surname)
	}
	fmt.Println("-------- Game --------")
	for _, game := range db.Games {
		fmt.Printf("%p\t| %v\t|%v\t| %v\t| %v\n", game.ID, game.Year, game.Category, game.Start, game.End)
	}

	fmt.Println("-------- Participation --------")
	for _, p := range db.Participations {
		fmt.Printf("%p\t| %p\t| %p\t| %p\t| %s\n", p.ID, p.Player, p.Game, p.Result, p.City)
	}
}
