package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

	var accounts = make(map[string]int)

	for _, tx := range records {

		valores, _ := strconv.Atoi(tx[1])

		if accounts[tx[0]] == 0 {
			accounts[tx[0]] = valores
		} else {
			accounts[tx[0]] = accounts[tx[0]] + valores
		}
	}
	for id, balance := range accounts {
		fmt.Printf("%v = %v \n", id, balance)
	}
}
