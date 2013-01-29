package main

import (
	"fmt"
	"github.com/earthboundkid/shuffle"
	"math/rand"
	"sort"
	"strconv"
)

const (
	MAX_SETS = 1
)

type Set sort.IntSlice

type Chromozome struct {
	sets  []*Set
	score int
}

func NewSet() *Set {
	var set Set
	shuffle.Shuffle(numbers)
	for i := 0; i < 6; i++ {
		set = append(set, numbers[i])
	}
	return &set
}

func (set *Set) getScore() int {
	score := 0
	matches := 0
	for _, row := range *test_data {
		matches = 0
		score += -1 //1 combination cost
		for _, num := range *set {
			n := sort.SearchInts(*row, num)
			if n < len(*row) && (*row)[n] == num {
				matches++
			}
		}
		switch matches {
		case 3:
			score += 1
		case 4:
			score += 2 * 2
		case 5:
			score += 2 * 3
		case 6:
			score += 2 * 4
		}
	}
	return score
}

func NewChromozome() *Chromozome {
	var ch Chromozome
	num_sets := 1
	if MAX_SETS > 1 {
		num_sets = rand.Intn(MAX_SETS) - 1
	}
	for i := 0; i < num_sets; i++ {
		ch.sets = append(ch.sets, NewSet())
	}
	ch.getScore()
	return &ch
}

func (ch *Chromozome) getScore() {
	ch.score = 0
	for _, set := range ch.sets {
		ch.score += set.getScore()
	}
}

func (ch *Chromozome) Print() {
	for _, set := range ch.sets {
		for _, num := range *set {
			fmt.Printf("%2s, ", strconv.Itoa(int(num)))
		}
		fmt.Printf("[%d]\n", ch.score)
	}
}
