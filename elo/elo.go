package elo

// Reelo returns the points for a given user calculated with a custom algorithm.
func Reelo(name, surname string) (reelo int) {
	// for y range years
	// if y isEmpty reelo = avg(pastFullYears)/2
	// if y isFull
	// cc = sum(Dmax)/Emax * Tmax
	// reelo = (cc * sum(D+E*K)/avgOfAvgCat) * 10000
	// 	if y is first full
	// 	reelo = (reelo + reelo/2)/2
	// reelo = reelo * Kaging
	return reelo
}
