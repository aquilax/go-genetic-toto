package main

import (
	"sort"
)

var numbers sort.IntSlice
var test_data *TestData

func main() {
	numbers = make(sort.IntSlice, 49, 49)
	for i := 1; i < 50; i++ {
		numbers[i-1] = i
	}
	ch := NewChromozome()
	ch.Print()
}
