package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
                [
                    {"a": 32, "b": 54},
                    {"a": 32, "b": 54},
                    {"a": 32, "b": 54},
                    {"a": 32, "b": 54},
                    {"a": 32, "b": 54}
                ]
            `
	type Message struct {
		A int32 `json:"a"`
		B int32 `json:"b"`
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.A, m.B)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

}
