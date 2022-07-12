package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func TestProfileMarshalUnmarshal(t *testing.T) {
	check.Check(t, check.Property(
		func(in Profile) error {
			data, err := json.Marshal(in)
			if err != nil {
				return fmt.Errorf("failed to encode: %w", err)
			}

			var out Profile
			if err := json.Unmarshal(data, &out); err != nil {
				return fmt.Errorf("failed to decode: %w", err)
			}

			if !reflect.DeepEqual(in, out) {
				return fmt.Errorf("encode/decode result doesn't match initial value")
			}

			return nil
		},
		generator.Struct(map[string]generator.Generator{
			"Name":    generator.String(),
			"Region":  generator.String(),
			"Raiting": generator.String(),
			"Interests": generator.Slice(generator.String()).Map(func(input []string) string {
				return strings.Join(input, ";")
			}),
		}),
	))
}
