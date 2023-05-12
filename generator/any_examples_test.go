package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Any() generator for generation of values for
// 3 types (int, uint and Point). Any() works for all go types except interfaces.
func ExampleAny() {
	type Point struct {
		X int16
		Y int16
		Z int8
	}

	streamer := generator.Streamer(
		func(i int, u uint, p Point) {
			fmt.Printf("%d, %d, %#v\n", i, u, p)
		},
		generator.Any(),
		generator.Any(),
		generator.Any(),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// -4518808235179270133, 7861855757602232086, generator_test.Point{X:5150, Y:-11651, Z:-40}
	// -2122761628320059770, 16172318933975836041, generator_test.Point{X:31913, Y:-20516, Z:-56}
	// -6485228379443441869, 9950499355085743487, generator_test.Point{X:31469, Y:16403, Z:-120}
	// 5972778420317109720, 18356510534673821416, generator_test.Point{X:-30288, Y:-3364, Z:-95}
	// 3277659493131553792, 16536373503259581585, generator_test.Point{X:-20183, Y:22302, Z:-117}
	// 2889002371220678169, 9050008079631751930, generator_test.Point{X:-2582, Y:-12674, Z:-7}
	// 1861954357100430827, 9301751449745155624, generator_test.Point{X:-16049, Y:12306, Z:74}
	// -1080827893950765636, 4332503251610791914, generator_test.Point{X:24459, Y:2740, Z:109}
	// -4610581452772180400, 11452572414278503431, generator_test.Point{X:3632, Y:-17110, Z:83}
	// 260101872073892018, 17022257214824873781, generator_test.Point{X:-5456, Y:-13741, Z:-12}
}
