package shrinker

import "github.com/steffnova/go-check/arbitrary"

func CollectionElements(original arbitrary.Arbitrary) arbitrary.Shrinker {
	transform := func(input arbitrary.Arbitrary) arbitrary.Arbitrary {
		input = input.Copy()
		for index, element := range original.Elements {
			input.Elements[index].Shrinker = element.Shrinker
		}

		return input
	}

	return Chain(
		CollectionOneElement().TransformOnceBefore(transform),
		CollectionAllElements().TransformOnceBefore(transform),
	)
}
