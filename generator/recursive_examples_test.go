package generator

import (
	"fmt"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func ExampleRecursive_binaryTree() {
	type Node struct {
		Value int
		Left  *Node
		Right *Node
	}

	var nodeString func(*Node) string
	nodeString = func(n *Node) string {
		if n == nil {
			return "nil"
		}
		return fmt.Sprintf("{Value: %d, Left: %s  Right %s}", n.Value, nodeString(n.Left), nodeString(n.Right))
	}

	err := Stream(0, 5, Streamer(
		func(tree *Node) {
			fmt.Println(nodeString(tree))
		},
		Recursive(func(r Recurse) arbitrary.Generator {
			return Ptr(Struct(map[string]arbitrary.Generator{
				"Value": Int(constraints.Int{Min: 0, Max: 10}),
				"Left":  r(),
				"Right": r(),
			}))
		}, 3),
	))

	if err != nil {
		panic(fmt.Errorf("Unexpected error: '%s'", err))
	}
	// Output:
	// {Value: 5, Left: nil  Right {Value: 8, Left: nil  Right nil}}
	// nil
	// {Value: 7, Left: nil  Right nil}
	// nil
	// {Value: 0, Left: nil  Right {Value: 7, Left: {Value: 3, Left: nil  Right nil}  Right nil}}
}

func ExampleRecursive_recursiveFunction() {
	type recursive func() (int, recursive)

	err := Stream(0, 10, Streamer(
		func(fn recursive) {
			ns := []int{}
			for fn != nil {
				var n int
				n, fn = fn()
				ns = append(ns, n)
			}
			fmt.Println(ns)
		},
		Recursive(func(r Recurse) arbitrary.Generator {
			return Weighted(
				[]uint64{4, 1},
				Func(Int(constraints.Int{Min: 0, Max: 10}), r()),
				Nil(),
			)
		}, 10),
	))

	if err != nil {
		panic(fmt.Errorf("Unexpected error: '%s'", err))
	}
	// Output:
	// [4 2 7 5 9 8 8]
	// [0]
	// [10]
	// [10 6]
	// [1 8 9]
	// []
	// [10 8 8 3 9 6 10 4 3 9 2]
	// [1 7]
	// []
	// []

}
