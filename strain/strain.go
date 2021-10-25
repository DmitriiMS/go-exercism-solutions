package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints *Ints) Keep(f func(int) bool) Ints {
	if *ints == nil {
		return nil
	}
	newInts := []int{}
	for _, i := range *ints {
		if f(i) {
			newInts = append(newInts, i)
		}
	}
	return newInts
}

func (ints *Ints) Discard(f func(int) bool) Ints {
	if *ints == nil {
		return nil
	}
	newInts := []int{}
	for _, i := range *ints {
		if !f(i) {
			newInts = append(newInts, i)
		}
	}
	return newInts
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
