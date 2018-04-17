package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("./data/transations.csv") // importando arquivo csv
	if err != nil {
		fmt.Print(err)
	}

	r := csv.NewReader(strings.NewReader(string(f)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("ReadAll():", err)
	}

	var accounts = make(map[string]string)

	for _, tx := range records {
		accounts[tx[0]] += tx[1]
	}

	for id, amount := range accounts {
		fmt.Printf("(%v, %v) \n", id, amount)
	}
}
