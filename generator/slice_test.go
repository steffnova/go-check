package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates how to use Slice(Int()) generator for generation of []int values.
// Slice requires generator for it's elements to be passed to it, thus Int() generator is used.
func ExampleSlice() {
	streamer := generator.Streamer(
		func(ints []int) {
			fmt.Printf("%#v\n", ints)
		},
		generator.Slice(
			generator.Int(constraints.Int{
				Min: 0,
				Max: 100,
			}),
		),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// []int{16, 80, 86, 69, 22, 84, 3, 64, 30, 91, 3, 45, 81, 31, 40, 62, 92, 21, 71, 50, 57, 76, 79, 95, 40, 69, 19, 94, 73, 9, 9}
	// []int{96, 81, 67, 36, 23, 59, 44, 52, 57, 56, 14, 30, 64, 77, 23, 29, 15, 52, 93, 51, 34, 1, 13, 44, 86, 68, 84, 40, 61, 19, 12, 57, 25, 83, 11, 0, 76, 87, 45, 73, 72}
	// []int{89, 12, 88, 96, 34, 10, 80, 74, 64, 27, 0, 36, 42, 7, 48, 4, 4, 95, 54, 46, 14, 43, 0, 17, 9, 17, 16, 13, 87, 19, 61, 30, 36, 38, 25, 47, 25, 56, 67, 88, 51, 29, 42, 62, 81}
	// []int{66, 22, 96, 2, 46, 21, 89, 12, 0, 7, 24, 10, 100, 38, 11, 18, 95, 28, 36, 100, 19, 69, 40, 94, 49, 89, 94, 30, 72, 48, 18, 6, 50, 28, 70, 79, 22, 74, 16, 68, 64, 70, 33, 42, 84, 11, 7, 3, 99, 52, 36, 46, 33, 64, 50, 32, 36, 48, 51, 81, 37, 36, 12, 7, 84, 48, 86, 64, 71, 86, 41, 41, 42, 47, 25, 17}
	// []int{37, 17, 94, 50, 16, 20, 53, 100, 80, 89, 93, 65, 9, 78, 16, 45, 75, 90, 54, 12, 24, 94, 60, 90, 37, 90, 80, 68, 98, 8, 91, 66, 45, 16, 49, 61, 52, 42, 62, 81, 67, 86, 21, 73, 61, 16, 3, 22, 27, 27, 11, 62, 24, 9, 64, 14, 84, 36, 35, 56, 62, 96, 15, 14, 38, 25, 10, 84, 42, 43, 93, 38, 12, 44, 43, 31, 12, 88, 66, 43, 47, 41, 40}
	// []int{70, 79, 28, 51, 70, 67, 16, 28, 97, 14, 87, 1, 91, 99, 24, 21, 23, 94, 39, 63, 59}
	// []int{68, 28, 45, 75, 4, 68, 42, 54, 80, 67, 53, 35, 84, 60, 2, 92, 28}
	// []int{45, 58, 33, 93, 14, 40, 80, 83, 67, 27, 21, 87, 25, 73, 50, 90, 75, 72, 62, 0, 32, 26, 56, 96, 6, 73, 39, 18, 84, 2, 95, 12, 85, 1, 43, 35, 77, 13, 6, 45, 4, 12, 34, 83, 59, 29, 1, 47, 69, 48, 100, 35, 42, 13, 19, 36, 67, 88, 24, 12, 1, 24, 80, 29}
	// []int{53, 30, 45, 39, 26, 98}
	// []int{26, 71, 84, 83, 33, 60, 15, 68, 98, 5, 97, 10, 27, 88, 86, 20, 54, 84, 5, 59, 19, 77, 14, 1, 97, 67, 0, 99, 12, 44, 99, 84, 95, 11, 13, 96, 67, 69, 55, 88, 68, 25, 47, 66, 98, 78, 57, 73, 14, 3, 51, 77, 39, 25, 90, 53, 74, 84, 22, 51, 79, 81, 93, 79, 43, 7, 48, 83, 29, 61, 29, 24, 68, 60, 64, 56, 10, 33, 97, 61, 70, 68, 93, 77, 1, 19, 92, 21, 44, 86}
}

// This example demonstrates how to use Slice(Int()) generator with constraints for generation of
// []int values. Constraints define range of generatable values for slice's size.
func ExampleSlice_constraints() {
	streamer := generator.Streamer(
		func(ints []int) {
			fmt.Printf("%#v\n", ints)
		},
		generator.Slice(
			generator.Int(constraints.Int{
				Min: 0,
				Max: 100,
			}),
			constraints.Length{Min: 0, Max: 10},
		),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// []int{31, 16, 80, 86, 69}
	// []int{84, 3, 64, 30, 91, 3}
	// []int{31}
	// []int{62, 92, 21, 71, 50, 57, 76, 79}
	// []int{95, 40}
	// []int{19, 94, 73, 9, 9}
	// []int{96, 81, 67, 36, 23, 59, 44, 52, 57}
	// []int{14, 30, 64, 77, 23, 29, 15, 52}
	// []int{34, 1, 13}
	// []int{68, 84, 40, 61, 19, 12}
}
