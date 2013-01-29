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
	set   *Set
	score int
}

func NewSet() *Set {
	var set Set
	shuffle.Shuffle(numbers)
	for i := 0; i < 6; i++ {
		set = append(set, numbers[i])
	}
	sort.Sort(sort.IntSlice(set))
	return &set
}

func (set *Set) getScore() int {
	score := 0
	matches := 0
	for _, row := range *test_data {
		matches = 0
		//		score += -1 //1 combination cost
		for _, num := range *set {
			n := sort.SearchInts(*row, num)
			if n < len(*row) && (*row)[n] == num {
				matches++
			}
		}
		switch matches {
		case 3:
			score += 2
		case 4:
			score += 2 << 1
		case 5:
			score += 2 << 2
		case 6:
			score += 2 << 3
		}
	}
	return score
}

func NewChromozome() *Chromozome {
	var ch Chromozome
	ch.set = NewSet()
	ch.score = ch.set.getScore()
	return &ch
}

func (ch *Chromozome) Print() {
	for _, num := range *ch.set {
		fmt.Printf("%2s, ", strconv.Itoa(int(num)))
	}
	fmt.Printf("[%d]\n", ch.score)
}

func (ch1 *Chromozome) dumbMate(ch2 *Chromozome) *Chromozome {
	var ch Chromozome
	var set Set
	length := len(*ch1.set)
	for i := 0; i < length; i++ {
		parent := rand.Intn(1)
		if parent == 0 {
			set = append(set, (*ch1.set)[i])
			continue
		}
		set = append(set, (*ch2.set)[i])

	}
	sort.Sort(sort.IntSlice(set))
	ch.set = &set
	ch.score = ch.set.getScore()
	return &ch
}
