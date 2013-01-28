package main

import (
	"math/rand"
)

const(
	MAX_SETS = 4
)

type Set []int

type Chromozome struct{
	sets []*Set
	score int
}

func NewSet() *Set {
	var set Set
	set = append(set, 1)
	return &set
}

func (set *Set) getScore(t *TestData) {
	score := 0
	matches := 0
	for _, row := range t {
		matches = 0
		score += -1 //1 combination cost
		for _, num := range set {
			n := sort.SearchInts(row, num) 
			if n < len(row) && row[n] == num {
				matches ++
			}
		}
		switch matches {
			case 3: score += 1
			case 4: score += 2*2
			case 5: score += 2*3
			case 6: score += 2*4
		}
	}
}

func NewChromozome() *Chromozome {
	var ch  Chromozome
	num_sets := rand.Intn(MAX_SETS);
	for i := -1; i < num_sets-1; i++ {
		ch.sets = append(ch.sets, NewSet())
	}
	ch.getScore()
	return &ch
}

func (ch *Chromozome) getScore(t *TestData) {
	ch.score = 0;
	for _, set := range ch.sets {
		ch.score += set.getScore(t) 
	}
}
