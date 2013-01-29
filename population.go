package main

import (
	"sort"
)

type Pool []*Chromozome

type Population struct {
	generation int
	pool       Pool
}

func (p Pool) Len() int           { return len(p) }
func (p Pool) Less(i, j int) bool { return p[i].score < p[j].score }
func (p Pool) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NewPopulation() *Population {
	return &Population{}
}

func (p *Population) Populate() {
	p.generation = 1
	for i := 0; i < options.pool_size; i++ {
		p.pool = append(p.pool, NewChromozome())
	}
	sort.Sort(p.pool)
}

func (p *Population) Crossover() *Population {
	p_new := NewPopulation()
	return p_new
}

func (p *Population) Print() {
	p.pool[0].Print()
}
