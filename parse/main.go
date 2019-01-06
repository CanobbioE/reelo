package main

func main() {
	// parse format of the file that is going to be parsed
	// TODO: This is an example for development purposes, this is the 2002 format.
	// Will need to get it in input somehow and iterate the years.
	inputFormat := []string{"Cognome", "Nome", "Città", "esercizi", "punti", "tempo"}
	format := newFormat(inputFormat)

	year := 2002
	categories := []string{"C1", "C2", "GP", "L1", "L2"}
	for _, category := range categories {

		// add the current year+category to the db
		//ctx := context.Background()
		//gID := db.Add(ctx, "giochi", year, category)

		// reads file and
		// parses the results and stores them in a DB
		readRankingFile(year, category, format)
	}
}
