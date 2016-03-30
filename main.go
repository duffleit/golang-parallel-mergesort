package main

import (
	"github.com/duffleit/mergesort/mergesort"
	"math/rand"
	"time"
	"log"
)


func main() {

	log.Print("This is a parallel mergesort written in go")
	log.Print("Run the included test to get see some statistics")

	numberOfItems := 1000000;
	threshold := 10000;

	items := rand.Perm(numberOfItems)

	start := time.Now()
	mergesort.MergeSort(items, threshold)

	log.Printf("Took %s to sort %d items (with threshold %d).", time.Since(start), numberOfItems, threshold)
}
