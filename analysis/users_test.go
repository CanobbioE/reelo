package analysis

import (
	"testing"

	"github.com/CanobbioE/reelo/parse"
)

func TestCompareUsers(t *testing.T) {
	var users = []struct {
		userA parse.User
		userB parse.User
		eOut  bool
	}{
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C1",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C1",
				Year:     2001,
			},
			true,
		},
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C1",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C1",
				Year:     2002,
			},
			false,
		},
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C1",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C2",
				Year:     2002,
			},
			true,
		},
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C2",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "GP",
				Year:     2007,
			},
			true,
		},
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C1",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "L1",
				Year:     2003,
			},
			true,
		},
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C1",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "L1",
				Year:     2004,
			},
			true,
		},
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "L2",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "GP",
				Year:     2001,
			},
			true,
		},
		{
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "C2",
				Year:     2000,
			},
			parse.User{
				Name:     "foo",
				Surname:  "bar",
				City:     "here",
				Category: "L1",
				Year:     2002,
			},
			true,
		},
	}

	for _, tt := range users {
		out := compareUsers(tt.userA, tt.userB)
		if out != tt.eOut {
			t.Errorf("Expected: %v for category %s, %d and %s, %d", tt.eOut, tt.userA.Category, tt.userA.Year, tt.userB.Category, tt.userB.Year)
		}
	}
}
