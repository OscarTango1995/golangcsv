package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"database/sql"
	"time"
)

func main() {
	readCSV("test.csv")

}

func readCSV(filename string){
	// Open the file
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			log.Fatalln("EOF reached but no data found! recheck that the file name is correct!")
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n",record[0], record[1],record[2], record[3],record[4], record[5],record[6], record[7],record[8], record[9],record[10], record[11],record[12], record[13],record[14], record[15],record[16], record[17],record[18], record[19],record[20])
	}

}
