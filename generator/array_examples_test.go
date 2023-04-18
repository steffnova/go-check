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
	// [5]int{-5339971465336467958, 5036824528102830934, 4435185786993720788, 8071137008395949086, 2122761628320059770}
	// [5]int{-5365688832259816617, -300681375570251064, -6485228379443441869, -8468275846115330281, -1089963290385541773}
	// [5]int{2727171422159354966, -315038161257240872, -660303368809814667, 5972778420317109720, 8502318506928285676}
	// [5]int{1284006505070580203, -2247583555303968036, 7505562437545936694, 710940327224637099, 58744246291326318}
	// [5]int{9177598355735269079, 4772086176229548406, -4788190396876772902, -3058895608739614131, 9050008079631751930}
	// [5]int{2430368660537815426, 5013668637975760780, -7056120859908864934, -8862094245172592907, 4700561838446243300}
	// [5]int{-5712649238675462878, 8746914360817110192, 325496396026881436, -3094703518683447370, -1080827893950765636}
	// [5]int{-4332503251610791914, 7551152765542507822, 5648976390688527282, 4610581452772180400, 3974191382882571532}
	// [5]int{2983335422402563632, 3236400634926555689, 260101872073892018, 8318806587974000740, -8405140618506968395}
	// [5]int{7656290481659077236, 1073273973791335576, -331846068917293018, 1135614103155420740, 7031127273457604653}
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
