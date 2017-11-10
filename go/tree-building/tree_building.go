package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Index struct {
	Record Record
	Node   *Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	indexes := make(map[int]Index)

	for _, record := range records {
		indexes[record.ID] = Index{Record: record, Node: &Node{ID: record.ID}}
	}

	if _, ok := indexes[0]; !ok {
		return nil, fmt.Errorf("no root node")
	}

	var keys []int
	for k := range indexes {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var tree *Node
	if root, ok := indexes[0]; ok {
		tree = root.Node
	} else {
		return nil, fmt.Errorf("tree does not have a root node")
	}

	for k := range keys {
		index := indexes[k]
		record := index.Record
		if record.ID != 0 {
			parent := indexes[record.Parent].Node
			if index.Record.ID == index.Record.Parent {
				return nil, fmt.Errorf("direct cyclical reference")
			}
			if parent.ID > index.Record.ID {
				return nil, fmt.Errorf("higher id parent of lower id")
			}
			parent.Children = append(parent.Children, index.Node)
		} else {
			if index.Record.Parent != 0 {
				return nil, fmt.Errorf("root node cannot have parent")
			}
			if index.Record.ID != k {
				return nil, fmt.Errorf("non-continous records")
			}
		}
	}

	return tree, nil
}

/*
Original
-----------------------------------------------------------------
BenchmarkTwoTree-8                     1        q ns/op
BenchmarkTenTree-8                    20          94844000 ns/op
BenchmarkShallowTree-8                 2         692778330 ns/op
PASS
ok      github.com/dotariel/exercism/go/tree-building   8.159s


Refactored
-----------------------------------------------------------------
BenchmarkTwoTree-8                    30          43085634 ns/op
BenchmarkTenTree-8                   300           4776996 ns/op
BenchmarkShallowTree-8               300           4737796 ns/op
PASS
ok      github.com/dotariel/exercism/go/tree-building   5.166s

*/
