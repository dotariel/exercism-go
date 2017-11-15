package strain

// Ints represents a slice of ints
type Ints []int

// Lists represents a slice of slice of ints
type Lists [][]int

// Strings represents a slice of strings
type Strings []string

// Keep returns a new Ints containing those elements where the
// predicate is true
func (l *Ints) Keep(f func(int) bool) Ints {
	if *l == nil {
		return nil
	}

	ints := Ints{}

	for _, it := range *l {
		if f(it) {
			ints = append(ints, it)
		}
	}

	return ints
}

// Discard returns a new Ints containing those elements where the
// predicate is false.
func (l *Ints) Discard(f func(int) bool) Ints {
	if *l == nil {
		return nil
	}

	ints := Ints{}

	for _, it := range *l {
		if !f(it) {
			ints = append(ints, it)
		}
	}

	return ints
}

// Keep returns a new Lists containing those elements where the
// predicate is true
func (ls *Lists) Keep(f func([]int) bool) Lists {
	if *ls == nil {
		return nil
	}

	lists := Lists{}

	for _, it := range *ls {
		if f(it) {
			lists = append(lists, it)
		}
	}

	return lists
}

// Keep returns a new Strings containing those elements where the
// predicate is true
func (ss *Strings) Keep(f func(string) bool) Strings {
	if *ss == nil {
		return nil
	}

	strings := Strings{}

	for _, it := range *ss {
		if f(it) {
			strings = append(strings, it)
		}
	}

	return strings
}
