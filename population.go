package main

import (
	"math/rand"
	"sort"
)

type Population []*Chromozome

func (p Population) Len() int           { return len(p) }
func (p Population) Less(i, j int) bool { return p[i].score > p[j].score }
func (p Population) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NewPopulation() *Population {
	return &Population{}
}

func (p *Population) Populate() {
	for i := 0; i < options.pool_size; i++ {
		*p = append(*p, NewChromozome())
	}
	sort.Sort(p)
}

func (p *Population) Crossover() *Population {
	p_new := NewPopulation()
	limit := (options.pool_size / 100) * 10
	for i := 0; i < limit; i++ {
		// Randomize sex
		r := rand.Intn(3) + 1
		*p_new = append(*p_new, (*p)[i].dumbMate((*p)[i+r]))
	}
	for j := limit; j < options.pool_size; j++ {
		*p_new = append(*p_new, NewChromozome())
	}
	sort.Sort(p_new)
	return p_new
}

func (p *Population) Print() {
	(*p)[0].Print()
}
