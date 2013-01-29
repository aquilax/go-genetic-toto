package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

const version string = "0.0.1"

var (
	options   Options
	numbers   sort.IntSlice
	test_data *TestData
)

func main() {
	processFlags()
	rand.Seed(time.Now().UTC().UnixNano())

	numbers = make(sort.IntSlice, 49, 49)
	for i := 1; i < 50; i++ {
		numbers[i-1] = i
	}
	test_data = NewTestData()
	if options.draws != "" {
		test_data.LoadFromFile(options.draws)
	} else {
		test_data.GenerateRandomData(options.num_draws)
	}

	population := NewPopulation()
	population.Populate()
	best := -99999999
	generation := 1
	for {
		population = population.Crossover()
		score := (*population)[0].score
		if score > best {
			fmt.Printf("%d\t", generation)
			population.Print()
			best = score
		}
		generation++
	}

	ch := NewChromozome()
	ch.Print()
}

func processFlags() {
	var fs = options.Init()
	fs.Parse(os.Args[1:])

	if options.version {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if options.help {
		fmt.Println("Genetic lotery number pcker: ", version)
		fmt.Println("Usage:")
		fs.PrintDefaults()
		os.Exit(0)
	}
}
