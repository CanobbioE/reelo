package reelo

var (
	startingYear           = 2002
	exercisesCostant       = 20.0
	pFinal                 = 1.5
	multiplicativeFactor   = 10000.0
	antiExploit            = 0.9
	noParticipationPenalty = 0.9
)

// Constants represent all the possible variables used when calculating the reelo
type Constants struct {
	StartingYear                                   int
	ExercisesCostant, PFinal, MultiplicativeFactor float64
	AntiExploit, NoParticipationPenalty            float64
}

// SlimArgs represents the arguments needed by the pseudoreelo function
type SlimArgs struct {
	IsParis                                      bool
	Start, End                                   int
	MaxScoreForCategory, Score, AvgCategoryScore float64
	Exercises                                    int
}

// Details represents a single year details for the given user
type Details struct {
	PseudoReelo    float64
	OldAvg, NewAvg float64
	Category       string
	// used to calculate max pseudo reelo
	Start, End                            int
	AvgCategoryScore, MaxScoreForCategory float64
}

// Args represents the arguments needed by the reelo function
type Args struct {
	Years                      []int
	LastKnownCategoryForPlayer string
	LastKnownYear              int
	History                    map[int]Details
}

// NewConstants istanciate a new set of costants
func NewConstants() Constants {
	return Constants{
		StartingYear:           2002,
		ExercisesCostant:       20.0,
		PFinal:                 1.5,
		MultiplicativeFactor:   10000.0,
		AntiExploit:            0.9,
		NoParticipationPenalty: 0.9,
	}
}

// InitConstants retrieves the costants in the database, if anything goes wrong
// it will fallback to the hardcoded values
// Variables names are chosen consistently with the formula
// provided by the scientific committee
func InitConstants(c Constants) {
	startingYear = c.StartingYear
	exercisesCostant = c.ExercisesCostant
	pFinal = c.PFinal
	multiplicativeFactor = c.MultiplicativeFactor
	antiExploit = c.AntiExploit
	noParticipationPenalty = c.NoParticipationPenalty
}

// CalculatePseudo calculates the player's Reelo given a single year's results.
// This means that the returned score does not take into account
// aging, anti-exploit and category promotion.
// The caller must provide all the arguments defined by the PseudoReeloArgs struct.
// The caller must handle the possibility a player has partecipated in two or more
// categories in the same year.
func CalculatePseudo(args SlimArgs) float64 {
	var baseScore float64
	t := args.Start
	n := args.End
	eMax := float64(n - t + 1)
	dMax := args.MaxScoreForCategory
	d := args.Score
	e := float64(args.Exercises)
	stepOne(&baseScore, e, d)
	stepTwo(&baseScore, args.IsParis)
	stepThree(&baseScore, t, n, d, e, eMax, dMax)
	stepFour(&baseScore, args.AvgCategoryScore)
	stepFive(&baseScore)

	return baseScore
}

// Calculate calculates a player's ELO using a custom algorithm
func Calculate(args Args) float64 {
	var reelo float64
	var sumOfWeights float64
	for _, year := range args.Years {
		curr := args.History[year]
		pseudoReelo := curr.PseudoReelo
		category := curr.Category
		oldAvg := curr.OldAvg
		newAvg := curr.NewAvg
		newMax := maxPseudoReelo(curr.Start, curr.End, curr.AvgCategoryScore, curr.MaxScoreForCategory)
		stepSix(&pseudoReelo, args.LastKnownCategoryForPlayer, category, year, newAvg, oldAvg, newMax)
		stepSeven(&pseudoReelo, &sumOfWeights, args.LastKnownYear, year)
		reelo += pseudoReelo
	}

	stepEight(&reelo, sumOfWeights)
	stepNine(&reelo, args.Years, args.LastKnownYear)
	stepTen(&reelo, args.Years, args.LastKnownYear)

	return reelo
}

// This is basically the funcion PseudoReelo with the obtained score is equal to
// the maximum score obtainable.
func maxPseudoReelo(start, end int, avgCategoryScore, maxScoreForCategory float64) float64 {
	var pseudoReelo float64
	t := start
	n := end
	eMax := float64(n - t + 1)
	e := eMax
	dMax := maxScoreForCategory
	d := dMax
	stepOne(&pseudoReelo, e, d)
	stepTwo(&pseudoReelo, true)
	stepThree(&pseudoReelo, t, n, d, e, eMax, dMax)
	stepFour(&pseudoReelo, avgCategoryScore)
	stepFive(&pseudoReelo)
	return pseudoReelo
}
