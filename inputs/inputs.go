package inputs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Categories   = []string{"C1", "C2", "L1", "L2", "GP"}
	StartingYear = 2002
	CurrentYear  = time.Now().Year()
	FolderPath   = "./ranks"
)

func DoStuff() {
	for y := StartingYear; y < CurrentYear; y++ {
		for cat := range Categories {
			filePath := fmt.Sprintf("%s/%d/%d_%s.txt", FolderPath, y, y, cat)
			makeDataStructure(filePath)
		}
	}
}

func makeDataStructure(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// do stuff with scanner.Bytes();
		// if word counts > 3 + scores fields chiedi all'utente
		// else salva in struttura dati
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
