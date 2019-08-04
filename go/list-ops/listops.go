package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (list IntList) Append(add IntList) IntList {
	for i := 0; i < len(add); i++ {
		list = append(list, add[i])
	}

	return list
}

func (list IntList) Concat(lists []IntList) IntList {

	for i := 0; i < len(lists); i++ {
		list = list.Append(lists[i])
	}

	return list
}

func (list IntList) Filter(fn predFunc) IntList {
	ret := make([]int, 0)

	for i := 0; i < len(list); i++ {
		if fn(list[i]) {
			ret = append(ret, list[i])
		}
	}

	return ret
}

func (list IntList) Foldr(fn binFunc, val int) int {
	for i := len(list) - 1; i >= 0; i-- {
		val = fn(list[i], val)
	}

	return val
}

func (list IntList) Foldl(fn binFunc, val int) int {
	for i := 0; i < len(list); i++ {
		val = fn(val, list[i])
	}

	return val
}

func (list IntList) Length() int {
	return len(list)
}

func (list IntList) Map(fn unaryFunc) IntList {
	ret := make([]int, 0)

	for i := 0; i < len(list); i++ {
		ret = append(ret, fn(list[i]))
	}

	return ret
}

func (list IntList) Reverse() IntList {
	ret := IntList(make([]int, 0))

	for i := (len(list) - 1); i >= 0; i-- {
		ret = append(ret, list[i])
	}

	return ret
}
