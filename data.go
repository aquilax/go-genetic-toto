package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	_ = iota
	ERR_INSUFFICIENT_NUMBERS
	ERR_NOT_A_NUMBER
	ERR_WRONG_NUMBER
	ERR_DUPLICATE_NUMBER
	ERR_FILE_OPEN_ERROR
)

const (
	MIN_NUMBER   = 1
	MAX_NUMBER   = 49
	DRAW_NUMBERS = 6
)

type TestData []*Set

func NewTestData() *TestData {
	return &TestData{}
}

func (t *TestData) GenerateRandomData(draws int) {
	for i := 0; i < draws; i++ {
		*t = append(*t, NewSet())
	}
}

func (t *TestData) LoadFromFile(file_name string) {
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Print(err)
		os.Exit(ERR_FILE_OPEN_ERROR)
	}

	defer f.Close()
	input := bufio.NewReader(f)
	for {
		line, err := input.ReadString('\n')
		if err == io.EOF {
			break
		}
		dash_ndx := strings.Index(line, "-")
		if dash_ndx < 0 {
			continue
		}
		draw_separator := strings.Index(line, "\t\t")
		*t = append(*t, getCombination(line[dash_ndx+1:draw_separator]))
		*t = append(*t, getCombination(line[draw_separator:]))
	}
}

func mytrim(s string) string {
	return strings.Trim(s, "\t \n")
}

func getCombination(numbers string) *Set {
	var s Set
	// strip space
	numbers = strings.Replace(mytrim(numbers), " ", "", -1)
	anumbers := strings.Split(numbers, ",")
	sort.Sort(sort.StringSlice(anumbers))
	if len(anumbers) != 6 {
		fmt.Println("ERROR: Please provide six integers separated by (,)")
		os.Exit(ERR_INSUFFICIENT_NUMBERS)
	}
	// store last number to easy get duplicates
	last_number := int64(-1)
	for _, snumber := range anumbers {
		inumber, e := strconv.ParseInt(snumber, 10, 8)
		if e != nil {
			fmt.Printf("ERROR: Wrong number: %s\n", snumber)
			os.Exit(ERR_NOT_A_NUMBER)
		}
		if inumber < MIN_NUMBER || inumber > MAX_NUMBER {
			fmt.Printf("ERROR: Please provide number in the range [%s..%d]. Provided: %s", MIN_NUMBER, MAX_NUMBER, snumber)
			os.Exit(ERR_WRONG_NUMBER)
		}
		if inumber == last_number {
			fmt.Println("ERROR: Duplicate number: " + snumber)
			os.Exit(ERR_DUPLICATE_NUMBER)
		}
		last_number = inumber
		s = append(s, int(inumber))
	}
	sort.Sort(sort.IntSlice(s))
	return &s
}
