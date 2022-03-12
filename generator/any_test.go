package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func ExampleAny() {
	type Point struct {
		X int16
		Y int16
		Z int8
	}

	// Streamer uses Any generator to generate int, uint and Point values.
	streamer := check.Streamer(
		func(i int, u uint, p Point) {
			fmt.Printf("%d, %d, %#v\n", i, u, p)
		},
		generator.Any(),
		generator.Any(),
		generator.Any(),
	)

	if err := check.Stream(streamer, check.Config{Iterations: 10, Seed: 0}); err != nil {
		panic(err)
	}

	// Output:
	// -5339971465336467958, 12088744466886928415, generator_test.Point{X:13142, Y:15828, Z:-91}
	// -1543285579645681342, 14677457169740829639, generator_test.Point{X:4175, Y:1247, Z:-116}
	// -300681375570251064, 5606570076237929230, generator_test.Point{X:-17391, Y:-25836, Z:-51}
	// -2023352169218621252, 9491810378858993108, generator_test.Point{X:21880, Y:23449, Z:120}
	// -7819249545370605693, 10732944964382368089, generator_test.Point{X:2272, Y:-30288, Z:-64}
	// -6787183051953194503, 15169603299902489319, generator_test.Point{X:-26463, Y:-21294, Z:0}
	// 9177598355735269079, 14220942032815928813, generator_test.Point{X:-26660, Y:17945, Z:25}
	// 9050008079631751930, 16728535719694244940, generator_test.Point{X:-870, Y:-12674, Z:-128}
	// -7056120859908864934, 1861954357100430827, generator_test.Point{X:15652, Y:-24979, Z:40}
	// -4265511144525599390, 11116133554876932735, generator_test.Point{X:-12306, Y:9628, Z:-74}
}
