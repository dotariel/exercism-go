package stringset

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Set []string

func New() Set {
	return Set{}
}

type ByKey []string

func (k ByKey) Len() int {
	return len(k)
}

func (k ByKey) Less(i, j int) bool {
	return k[i] < k[j]
}

func (k ByKey) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func NewFromSlice(sl []string) Set {
	res := Set{}

	for _, val := range sl {
		res.Add(val)
	}

	return res
}

func (s Set) String() string {
	f := func(s string) string {
		return fmt.Sprintf(`"%v"`, s)
	}

	return fmt.Sprintf(`{%v}`, strings.Join(s.Map(f), ", "))
}

func (s Set) Map(f func(string) string) Set {
	tmp := make([]string, len(s))

	for i, val := range s {
		tmp[i] = f(val)
	}

	return Set(tmp)
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(item string) bool {
	for _, val := range s {
		if item == val {
			return true
		}
	}

	return false
}

func Subset(s1, s2 Set) bool {
	if s1.IsEmpty() || Equal(s1, s2) {
		return true
	}

	found := 0
	for _, item := range s1 {
		if s2.Has(item) {
			found++
		}
	}

	return found == len(s1)
}

func Disjoint(s1, s2 Set) bool {
	if Equal(s1, s2) {
		return true
	}

	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}

	return len(Intersection(s1, s2)) == 0
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func (s *Set) Add(item string) {
	if !s.Has(item) {
		*s = append(*s, item)
	}
	sort.Sort(ByKey(*s))
}

func Intersection(s1, s2 Set) Set {
	res := Set{}

	for _, item := range s1 {
		if s2.Has(item) {
			res = append(res, item)
		}
	}

	return res
}

func Difference(s1, s2 Set) Set {
	res := Set{}

	if s1.IsEmpty() {
		return res
	}

	if s2.IsEmpty() {
		return s1
	}

	for _, item := range s1 {
		if !s2.Has(item) {
			res = append(res, item)
		}
	}

	sort.Sort(ByKey(res))

	return res
}

func Union(s1, s2 Set) Set {
	res := Set{}

	for _, item := range s1 {
		res.Add(item)
	}

	for _, item := range s2 {
		res.Add(item)
	}

	return res
}
