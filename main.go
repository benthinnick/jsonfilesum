package main

import (
	"encoding/json"
	"fmt"
	"gofilesum/types"
	"log"
	"os"
	"sync"
)

const bulkSize = 10

func startProcess(in <-chan [bulkSize]types.Pair, out chan<- int64, wg *sync.WaitGroup) {

	go func() {
		defer wg.Done()
		var sum int64

		for pairs := range in {
			for _, pair := range pairs {
				sum += (pair.A + pair.B)
			}
			out <- sum
			sum = 0
		}

		log.Println("Channel closed")

		out <- sum
	}()
}

func main() {
	var wg sync.WaitGroup

	var chans [5]chan [bulkSize]types.Pair
	out := make(chan int64)
	res := make(chan int64)

	for i := 0; i < 5; i++ {
		chans[i] = make(chan [bulkSize]types.Pair, 3)
		wg.Add(1)
		startProcess(chans[i], out, &wg)
	}

	go func() {
		var s int64
		for v := range out {
			s += v
		}

		res <- s
	}()

	f, err := os.Open("./bigData.json")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	var pairs [bulkSize]types.Pair
	currChan := 0
	for dec.More() {
		var m types.Pair
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		pairs[count] = m
		count++
		if count >= bulkSize {
			chans[currChan] <- pairs
			currChan++
			if currChan >= len(chans) {
				currChan = 0
			}
			count = 0
			pairs = [bulkSize]types.Pair{}
		}
	}
	chans[currChan] <- pairs

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for _, ch := range chans {
		close(ch)
	}

	wg.Wait()
	log.Println("Channels closed")
	close(out)

	ans, ok := <-res
	if ok {
		fmt.Println(ans)
	}

}
