package parse

// Limits represents the starting end ending exercise for each category
type Limits struct {
	Start, End int
}

var formats = map[int]string{
	2003: "cognome nome esercizi punti tempo",
	2004: "cognome nome esercizi punti città",
	2005: "cognome nome città esercizi punti",
	2006: "cognome nome città esercizi punti",
	2007: "cognome nome città esercizi punti",
	2008: "cognome nome città esercizi punti",
	2009: "cognome nome città esercizi punti",
	2010: "cognome nome città esercizi punti",
	2011: "cognome nome città esercizi punti",
	2012: "cognome nome città esercizi punti",
	2013: "cognome nome città esercizi punti",
	2014: "cognome nome città esercizi punti",
	2015: "cognome nome città esercizi punti",
	2016: "cognome nome città esercizi punti",
	2017: "cognome nome città esercizi punti",
	2018: "cognome nome città esercizi punti",
}

var scores2002 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   8,
	},
	"C2": {
		Start: 3,
		End:   10,
	},
	"L1": {
		Start: 5,
		End:   12,
	},
	"L2": {
		Start: 6,
		End:   14,
	},
	"GP": {
		Start: 6,
		End:   14,
	},
}

var scores2003 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   8,
	},
	"C2": {
		Start: 3,
		End:   10,
	},
	"L1": {
		Start: 5,
		End:   12,
	},
	"L2": {
		Start: 6,
		End:   14,
	},
	"GP": {
		Start: 5,
		End:   14,
	},
}

var scores2004 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   8,
	},
	"C2": {
		Start: 3,
		End:   10,
	},
	"L1": {
		Start: 4,
		End:   11,
	},
	"L2": {
		Start: 6,
		End:   13,
	},
	"GP": {
		Start: 7,
		End:   14,
	},
}

var scores2005 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   8,
	},
	"C2": {
		Start: 3,
		End:   10,
	},
	"L1": {
		Start: 5,
		End:   12,
	},
	"L2": {
		Start: 6,
		End:   13,
	},
	"GP": {
		Start: 7,
		End:   14,
	},
}

var scores2006 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   9,
	},
	"C2": {
		Start: 3,
		End:   10,
	},
	"L1": {
		Start: 5,
		End:   12,
	},
	"L2": {
		Start: 7,
		End:   14,
	},
	"GP": {
		Start: 8,
		End:   15,
	},
}

var scores2007 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 5,
		End:   12,
	},
	"L1": {
		Start: 7,
		End:   14,
	},
	"L2": {
		Start: 9,
		End:   16,
	},
	"GP": {
		Start: 9,
		End:   18,
	},
}

var scores2008 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 4,
		End:   12,
	},
	"L1": {
		Start: 7,
		End:   14,
	},
	"L2": {
		Start: 9,
		End:   16,
	},
	"GP": {
		Start: 10,
		End:   18,
	},
}

var scores2009 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 3,
		End:   12,
	},
	"L1": {
		Start: 5,
		End:   14,
	},
	"L2": {
		Start: 7,
		End:   16,
	},
	"GP": {
		Start: 8,
		End:   18,
	},
}

var scores2010 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 3,
		End:   12,
	},
	"L1": {
		Start: 5,
		End:   14,
	},
	"L2": {
		Start: 7,
		End:   16,
	},
	"GP": {
		Start: 1,
		End:   16,
	},
}

var scores2011 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   9,
	},
	"C2": {
		Start: 1,
		End:   10,
	},
	"L1": {
		Start: 3,
		End:   12,
	},
	"L2": {
		Start: 5,
		End:   14,
	},
	"GP": {
		Start: 5,
		End:   16,
	},
}

var scores2012 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 1,
		End:   12,
	},
	"L1": {
		Start: 8,
		End:   16,
	},
	"L2": {
		Start: 10,
		End:   18,
	},
	"GP": {
		Start: 9,
		End:   20,
	},
}

var scores2013 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   9,
	},
	"C2": {
		Start: 2,
		End:   11,
	},
	"L1": {
		Start: 4,
		End:   13,
	},
	"L2": {
		Start: 7,
		End:   16,
	},
	"GP": {
		Start: 1,
		End:   18,
	},
}

var scores2014 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 3,
		End:   12,
	},
	"L1": {
		Start: 5,
		End:   14,
	},
	"L2": {
		Start: 7,
		End:   16,
	},
	"GP": {
		Start: 1,
		End:   18,
	},
}

var scores2015 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   9,
	},
	"C2": {
		Start: 1,
		End:   12,
	},
	"L1": {
		Start: 4,
		End:   14,
	},
	"L2": {
		Start: 6,
		End:   16,
	},
	"GP": {
		Start: 7,
		End:   18,
	},
}

var scores2016 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 1,
		End:   12,
	},
	"L1": {
		Start: 3,
		End:   14,
	},
	"L2": {
		Start: 5,
		End:   16,
	},
	"GP": {
		Start: 1,
		End:   17,
	},
}

var scores2017 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 1,
		End:   12,
	},
	"L1": {
		Start: 3,
		End:   14,
	},
	"L2": {
		Start: 5,
		End:   16,
	},
	"GP": {
		Start: 1,
		End:   18,
	},
}

var scores2018 = map[string]Limits{
	"C1": {
		Start: 1,
		End:   10,
	},
	"C2": {
		Start: 1,
		End:   12,
	},
	"L1": {
		Start: 3,
		End:   14,
	},
	"L2": {
		Start: 5,
		End:   16,
	},
	"GP": {
		Start: 1,
		End:   18,
	},
}

var scoreHelp = map[int]map[string]Limits{
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
