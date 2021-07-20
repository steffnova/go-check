package generator

import (
	"github.com/steffnova/go-check/arbitrary"
	"math/rand"
	"reflect"
)

// Generate, generates random arbitrary.Type using rand
type Generate func(rand *rand.Rand) arbitrary.Type

// Type is a type generator. Type parameter defines underlying type of values
// generated with Generate.
type Type struct {
	Generate
	reflect.Type
}
