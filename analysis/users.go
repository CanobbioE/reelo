package analysis

import (
	"fmt"

	"github.com/CanobbioE/reelo/parse"
)

func agglomerate() {
	users := parse.GetUsers()
	for index, user := range users {
		fmt.Println(index, user)
	}
}

func compareUsers(a, b parse.User) bool {
	if a.Name == b.Name &&
		a.Surname == b.Surname &&
		a.City == b.City { // Not sure if we want to check the city as well...
		if a.Year < b.Year {
			switch a.Category {
			case "C1":
				switch b.Category {
				case "C1":
					if b.Year-a.Year == 1 {
						return true
					}
				case "C2":
					if b.Year-a.Year <= 3 {
						return true
					}
				case "L1":
					if b.Year-a.Year <= 6 {
						return true
					}
				case "L2":
					if b.Year-a.Year <= 9 {
						return true
					}
				case "GP":
					if b.Year-a.Year >= 9 {
						return true
					}
				}
			case "C2":
				switch b.Category {
				case "C1":
					return false
				case "C2":
					if b.Year-a.Year == 1 {
						return true
					}
				case "L1":
					if b.Year-a.Year <= 4 {
						return true
					}
				case "L2":
					if b.Year-a.Year <= 7 {
						return true
					}
				case "GP":
					if b.Year-a.Year >= 7 {
						return true
					}
				}
			case "L1":
				switch b.Category {
				case "C1":
					return false
				case "C2":
					return false
				case "L1":
					if b.Year-a.Year <= 2 {
						return true
					}
				case "L2":
					if b.Year-a.Year <= 5 {
						return true
					}
				case "GP":
					if b.Year-a.Year >= 4 {
						return true
					}
				}
			case "L2":
				switch b.Category {
				case "C1":
					return false
				case "C2":
					return false
				case "L1":
					return false
				case "L2":
					if b.Year-a.Year <= 2 {
						return true
					}
				case "GP":
					if b.Year-a.Year >= 1 {
						return true
					}
				}
			case "GP":
				switch b.Category {
				case "C1":
					return false
				case "C2":
					return false
				case "L1":
					return false
				case "L2":
					return false
				case "GP":
					if b.Year-a.Year >= 1 {
						return true
					}
				}
			}
		}
	}
	return false
}

/*
C1 -> 11, 12 yrs
C2 -> 13, 14 yrs
L1 -> 15, 16, 17 yrs
L2 -> 18, 19, 20 yrs
GP -> 21, ... 99 yrs
*/
