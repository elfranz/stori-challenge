package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// define a transaction
type Transaction struct {
	Date     string
	Amount   float64
	IsCredit bool
}

func main() {
	// File path and email details
	csvPath := "transactions.csv" // should be env variable since its not gonna change
	// recipientEmail := "user@example.com" // placeholder, should be env variable

	// process csv
	transactions, err := readTransactions(csvPath)
	if err != nil {
		log.Fatalf("Failed to process transactions: %v", err)
	}
	fmt.Printf("%v", transactions)
	// calculate summary

	// format email

	// send email

}

func readTransactions(csvPath string) ([]Transaction, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	var transactions []Transaction

	for {
		record, err := reader.Read()
		if err != nil {
			break // EOF
		}

		// im having trouble parsing the date so for now I will just add the string as is, need to add validations and parsing to not overcomplicate calculations later on
		// date, err := time.Parse("%m,%d", record[1])
		// if err != nil {
		// 	return nil, err
		// }

		date := record[1]

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, Transaction{
			Date:     date,
			Amount:   amount,
			IsCredit: amount > 0,
		})
	}

	return transactions, nil
}
