package shrinker

import "github.com/steffnova/go-check/arbitrary"

func Fail(err error) arbitrary.Shrinker {
	return arbitrary.Shrinker(nil).Fail(err)
}
