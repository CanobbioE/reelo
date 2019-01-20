package elo

type exercises struct {
	start int
	end   int
}

//TODO: "Scores" should be something like "Exercises each category has to complete"
type Scores map[string]exercises

// StartOfCategory returns the number of the first exercise for the specified year and category
func StartOfCategory(year int, category string) int {
	return allScores[year][category].start
}

// EndOfCategory returns the number of the last exercise for the specified year and category
func EndOfCategory(year int, category string) int {
	return allScores[year][category].end
}

// I don't really like this but...
// This grants clear access to data
var allScores = map[int]Scores{
	2002: scores2002,
	2003: scores2003,
	2004: scores2004,
	2005: scores2005,
	2006: scores2006,
	2007: scores2007,
	2008: scores2008,
	2009: scores2009,
	2010: scores2010,
	2011: scores2011,
	2012: scores2012,
	2013: scores2013,
	2014: scores2014,
	2015: scores2015,
	2016: scores2016,
	2017: scores2017,
	2018: scores2018,
}

// TODO: double-check all the start/end values

var scores2002 = Scores{
	"C1": exercises{
		start: 1,
		end:   8,
	},
	"C2": exercises{
		start: 3,
		end:   10,
	},
	"L1": exercises{
		start: 5,
		end:   12,
	},
	"L2": exercises{
		start: 6,
		end:   14,
	},
	"GP": exercises{
		start: 6,
		end:   14,
	},
}

var scores2003 = Scores{
	"C1": exercises{
		start: 1,
		end:   8,
	},
	"C2": exercises{
		start: 3,
		end:   10,
	},
	"L1": exercises{
		start: 5,
		end:   12,
	},
	"L2": exercises{
		start: 6,
		end:   14,
	},
	"GP": exercises{
		start: 5,
		end:   14,
	},
}

var scores2004 = Scores{
	"C1": exercises{
		start: 1,
		end:   8,
	},
	"C2": exercises{
		start: 3,
		end:   10,
	},
	"L1": exercises{
		start: 4,
		end:   11,
	},
	"L2": exercises{
		start: 6,
		end:   13,
	},
	"GP": exercises{
		start: 7,
		end:   14,
	},
}

var scores2005 = Scores{
	"C1": exercises{
		start: 1,
		end:   8,
	},
	"C2": exercises{
		start: 3,
		end:   10,
	},
	"L1": exercises{
		start: 5,
		end:   12,
	},
	"L2": exercises{
		start: 6,
		end:   13,
	},
	"GP": exercises{
		start: 7,
		end:   14,
	},
}

var scores2006 = Scores{
	"C1": exercises{
		start: 1,
		end:   9,
	},
	"C2": exercises{
		start: 3,
		end:   10,
	},
	"L1": exercises{
		start: 5,
		end:   12,
	},
	"L2": exercises{
		start: 7,
		end:   14,
	},
	"GP": exercises{
		start: 8,
		end:   15,
	},
}

var scores2007 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 5,
		end:   12,
	},
	"L1": exercises{
		start: 7,
		end:   14,
	},
	"L2": exercises{
		start: 9,
		end:   16,
	},
	"GP": exercises{
		start: 9,
		end:   18,
	},
}

var scores2008 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 4,
		end:   12,
	},
	"L1": exercises{
		start: 7,
		end:   14,
	},
	"L2": exercises{
		start: 9,
		end:   16,
	},
	"GP": exercises{
		start: 10,
		end:   18,
	},
}

var scores2009 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 3,
		end:   12,
	},
	"L1": exercises{
		start: 5,
		end:   14,
	},
	"L2": exercises{
		start: 7,
		end:   16,
	},
	"GP": exercises{
		start: 8,
		end:   18,
	},
}

var scores2010 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 3,
		end:   12,
	},
	"L1": exercises{
		start: 5,
		end:   14,
	},
	"L2": exercises{
		start: 7,
		end:   16,
	},
	"GP": exercises{
		start: 1,
		end:   16,
	},
}

var scores2011 = Scores{
	"C1": exercises{
		start: 1,
		end:   9,
	},
	"C2": exercises{
		start: 1,
		end:   10,
	},
	"L1": exercises{
		start: 3,
		end:   12,
	},
	"L2": exercises{
		start: 5,
		end:   14,
	},
	"GP": exercises{
		start: 5,
		end:   16,
	},
}

var scores2012 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 1,
		end:   12,
	},
	"L1": exercises{
		start: 8,
		end:   16,
	},
	"L2": exercises{
		start: 10,
		end:   18,
	},
	"GP": exercises{
		start: 9,
		end:   20,
	},
}

var scores2013 = Scores{
	"C1": exercises{
		start: 1,
		end:   9,
	},
	"C2": exercises{
		start: 2,
		end:   11,
	},
	"L1": exercises{
		start: 4,
		end:   13,
	},
	"L2": exercises{
		start: 7,
		end:   16,
	},
	"GP": exercises{
		start: 1,
		end:   18,
	},
}

var scores2014 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 3,
		end:   12,
	},
	"L1": exercises{
		start: 5,
		end:   14,
	},
	"L2": exercises{
		start: 7,
		end:   16,
	},
	"GP": exercises{
		start: 1,
		end:   18,
	},
}

var scores2015 = Scores{
	"C1": exercises{
		start: 1,
		end:   9,
	},
	"C2": exercises{
		start: 1,
		end:   12,
	},
	"L1": exercises{
		start: 4,
		end:   14,
	},
	"L2": exercises{
		start: 6,
		end:   16,
	},
	"GP": exercises{
		start: 7,
		end:   18,
	},
}

var scores2016 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 1,
		end:   12,
	},
	"L1": exercises{
		start: 3,
		end:   14,
	},
	"L2": exercises{
		start: 5,
		end:   16,
	},
	"GP": exercises{
		start: 1,
		end:   17,
	},
}

var scores2017 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 1,
		end:   12,
	},
	"L1": exercises{
		start: 3,
		end:   14,
	},
	"L2": exercises{
		start: 5,
		end:   16,
	},
	"GP": exercises{
		start: 1,
		end:   18,
	},
}

var scores2018 = Scores{
	"C1": exercises{
		start: 1,
		end:   10,
	},
	"C2": exercises{
		start: 1,
		end:   12,
	},
	"L1": exercises{
		start: 3,
		end:   14,
	},
	"L2": exercises{
		start: 5,
		end:   16,
	},
	"GP": exercises{
		start: 1,
		end:   18,
	},
}
