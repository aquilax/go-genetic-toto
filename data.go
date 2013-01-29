package main

type TestData []*Set

func NewTestData() *TestData {
	return &TestData{}
}

func (t *TestData) LoadFromFile(file_name string) {
	//FIXME dummy
	for i := 0; i < 6; i++ {
		set := &Set{1, 2, 3, 4, 5, 6}
		*t = append(*t, set)
	}
}
