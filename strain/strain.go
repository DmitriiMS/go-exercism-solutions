package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints *Ints) Keep(f func(int) bool) (newInts Ints) { //named returns are nil until modified, so this varian skips nil checks
	for _, i := range *ints {
		if f(i) {
			newInts = append(newInts, i)
		}
	}
	return newInts
}

func (ints *Ints) Discard(f func(int) bool) (newInts Ints) {
	return ints.Keep(func(i int) bool { return !f(i) }) //use reverse fucntion in a wrapper function
}

func (lists *Lists) Keep(f func([]int) bool) Lists {
	newLists := [][]int{}
	for _, l := range *lists {
		if f(l) {
			newLists = append(newLists, l)
		}
	}
	return newLists
}

func (strings *Strings) Keep(f func(string) bool) Strings {
	newStrings := []string{}
	for _, s := range *strings {
		if f(s) {
			newStrings = append(newStrings, s)
		}
	}
	return newStrings
}
