package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Array(Int()) generator for generation of
// array of integers with 5 elements. Array() requires a generator for it's
// elements to be passed to it. Int() generator is used as an element generator.
func ExampleArray() {
	streamer := generator.Streamer(
		func(arr [5]int) {
			fmt.Printf("%#v\n", arr)
		},
		generator.Array(generator.Int()),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// [5]int{-4518808235179270133, -5036824528102830934, 8071137008395949086, -2122761628320059770, 6925466992496964832}
	// [5]int{5606570076237929230, -6485228379443441869, -1089963290385541773, -315038161257240872, -8800248522230157011}
	// [5]int{5309191430948421708, 8901768139528370850, -1284006505070580203, 3029607914297333936, -7505562437545936694}
	// [5]int{9177598355735269079, 2353700405777826677, 2889002371220678169, 2430368660537815426, -5013668637975760780}
	// [5]int{1861954357100430827, 3930565546038334149, 4265511144525599390, 9001397090059497490, -7811778195032818482}
	// [5]int{3434633470139019496, 4491726675100777446, 7551152765542507822, 6992623394542015136, -4610581452772180400}
	// [5]int{2983335422402563632, -3236400634926555689, -8318806587974000740, 7656290481659077236, -1073273973791335576}
	// [5]int{1135614103155420740, -7031127273457604653, 2574016966663937194, 5790223617635911539, 8164215293865182704}
	// [5]int{-6776706801480370880, -2849222499850869647, -8212568795191619498, -1767412366155247756, -6821913011751915842}
	// [5]int{7534629738184250638, 1276965762818039701, -1600752837639152423, -6612685186773303684, -4142208901106344656}
}

// This example demonstrates usage of ArrayFrom() generator and Int() generator for generation
// of array of integers with 5 elements. ArrayFrom() requires generator for each element of the
// array to be passed to it. Five Int() generator are used for array elements.
func ExampleArrayFrom() {
	streamer := generator.Streamer(
		func(arr [5]int) {
			fmt.Printf("%#v\n", arr)
		},
		generator.ArrayFrom(
			generator.Int(constraints.Int{Min: 0, Max: 9}),
			generator.Int(constraints.Int{Min: 10, Max: 19}),
			generator.Int(constraints.Int{Min: 20, Max: 29}),
			generator.Int(constraints.Int{Min: 30, Max: 39}),
			generator.Int(constraints.Int{Min: 40, Max: 49}),
		),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// [5]int{5, 12, 26, 30, 40}
	// [5]int{6, 11, 25, 36, 44}
	// [5]int{3, 10, 23, 31, 48}
	// [5]int{8, 15, 27, 32, 49}
	// [5]int{2, 18, 25, 33, 49}
	// [5]int{9, 14, 29, 39, 40}
	// [5]int{1, 13, 24, 37, 44}
	// [5]int{9, 18, 20, 36, 49}
	// [5]int{7, 14, 23, 32, 41}
	// [5]int{6, 14, 24, 38, 48}
}
