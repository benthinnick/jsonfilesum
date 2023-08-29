package main

import (
	"encoding/json"
	"fmt"
	"gofilesum/types"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./bigData.json")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var sumA int64
	var sumB int64
	dec := json.NewDecoder(f)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var m types.Pair
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		sumA += m.A
		sumB += m.B
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sumA: %v\n", sumA)
	fmt.Printf("sumB: %v\n", sumB)

}
