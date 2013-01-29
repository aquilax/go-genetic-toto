package main

import (
	"math/rand"
	"sort"
	"time"
)

var numbers sort.IntSlice
var test_data *TestData

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	numbers = make(sort.IntSlice, 49, 49)
	for i := 1; i < 50; i++ {
		numbers[i-1] = i
	}
	test_data = NewTestData()
	test_data.LoadFromFile("test")
	ch := NewChromozome()
	ch.Print()
}
