package generator

import (
	"errors"
	"testing"

	"github.com/steffnova/go-check/arbitrary"
	"github.com/steffnova/go-check/constraints"
)

func TestRecursive(t *testing.T) {
	type Node struct {
		Value int
		Left  *Node
		Right *Node
	}

	testCase := map[string]func(*testing.T){
		"InvalidTarget": func(t *testing.T) {
			err := Stream(0, 10, Streamer(
				func(int) {},
				Recursive(func(r Recurse) arbitrary.Generator {
					return Struct(map[string]arbitrary.Generator{
						"Value": Int(constraints.Int{Min: 0, Max: 10}),
						"Left":  r(),
						"Right": r(),
					})
				}, 0),
			))

			if !errors.Is(err, arbitrary.ErrorInvalidTarget) {
				t.Fatalf("Expected error: '%s'", arbitrary.ErrorInvalidTarget)
			}
		},
		"Depth0": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(tree *Node) {
					if tree.Left != nil || tree.Right != nil {
						t.Fatalf("Tree shouldn't have left and right nodes if depth is 0")
					}
				},
				Recursive(func(r Recurse) arbitrary.Generator {
					return PtrTo(Struct(map[string]arbitrary.Generator{
						"Value": Int(constraints.Int{Min: 0, Max: 10}),
						"Left":  r(),
						"Right": r(),
					}))
				}, 0),
			))

			if err != nil {
				t.Fatalf("Unexpected error: '%s'", err)
			}
		},
		"Depth5": func(t *testing.T) {
			err := Stream(0, 100, Streamer(
				func(tree *Node) {
					height := (func(node *Node) int)(nil)
					height = func(node *Node) int {
						if node == nil {
							return 0
						}
						left := height(node.Left)
						right := height(node.Right)

						if left > right {
							return 1 + left
						}
						return 1 + right
					}
					h := height(tree)
					if h != 6 {
						t.Fatalf("Tree should have a height of 6, got %d", h)
					}
				},
				Recursive(func(r Recurse) arbitrary.Generator {
					return PtrTo(Struct(map[string]arbitrary.Generator{
						"Value": Int(constraints.Int{Min: 0, Max: 10}),
						"Left":  r(),
						"Right": r(),
					}))
				}, 5),
			))

			if err != nil {
				t.Fatalf("Unexpected error: '%s'", err)
			}
		},
	}

	for name, testCase := range testCase {
		t.Run(name, testCase)
	}
}
