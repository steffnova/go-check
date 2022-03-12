# Generators

In `go-check` generators are fundamental building blocks for writing test. Generators are used to generated data of specified type. They define constraints that control how data is generated and for generated data they usually provide corresponding shrinker. Goal of generators is to be a powerful tool for generating any type of data. 

There are two ways of defining generators:
  - Implementing one from scratch
  - Building new generator with Combinators 

# Combinators

Combinators allow manipulation of generated data, which consists of adding new constraints (Filter), mapping generated data (Map), or using generated data for an input to another generator (Bind). Combinators can be used in any order and any number of times thus making them a powerful tool for expressing constraints and structure of data. All combinators return a derived Generator with altered behavior from original.

  - [Map](#map)
  - [Filter](#filter)
  - [Bind](#bind)

## Map

Map combinator maps generated data using mapper. Mapper is any function that has one input and one output. Mapper's input must be of the same type as the type of data being generated by it's source. Mapper's output type defines type of data being generated by derived generator. Following example shows how Uint generator along with Map combinator can be used to define generator for bool type:

```go
package main_test

import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/constraints"
    "github.com/steffnova/go-check/generator"
)

func TestBool(t *testing.T) {
    // Source generator is Uint generator that can generate either 0 or 1.
    // 0 and 1 constraints are imposed by passing constraints.Uint to Uint
    // generator.
    source := generator.Uint(constraints.Uint{
        Min: 0,
        Max: 1,
    })

    // Mapper function will take uint as an input and return bool as result.
    mapper := func(n uint) bool {
        return n == 1
    }

    // Bool generator is defined using Map combinator
    derived := source.Map(mapper)

    // Derived generator can now used by Check for generating bool values
    check.Check(t, check.Property(
        func(b bool) error {
            return nil
        },
        derived,
    ))
}
```

Map generator can be used to alter generated data. This feature is extremly helpful when there is a need to introduce structure in random data. One good example is generating even integer numbers:


```go
package main_test

import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/constraints"
    "github.com/steffnova/go-check/generator"
)

func TestEvenInteger(t *testing.T) {
    check.Check(t, check.Property(
        func(n int) error {
            if n%2 != 0 {
                return fmt.Errorf("Not an even integer: %d", n)
            }
            return nil
        },
        generator.Int().Map(func(n int) int {
            return n*2
        }),
    ))
}
```

## Filter
Filter combinator filters generated data using predicate. Predicate is a function that has one input and one output. Predicate's input must be of the same type as the type of data being generated by it's source and it's output is always bool. If data generated by the source doesn't satisfy predicate, generator will generate new sample. This process will repeat until predicate is satisfied. Next example demonstrates how Filter combinator can be used to implement generator for even integers unlike the unlike the previous one where Map combinator was used:

```go
package main_test

import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/generator"
)

func TestEvenInteger(t *testing.T) {
    // Predicate will filter even integer numbers.
    predicate := func(n uint) bool {
        return n%2 == 0
    }

    check.Check(t, check.Property(
        func(n int) error {
            return nil
        },

        generator.Int().Filter(predicate),
    ))
}
```

*NOTE*: Filter combinator can impact generator's generation speed, or in the worst case enter a infinite loop, so a special care needs to be attended when choosing predicate for generator's data set. Two generators used in following example demonstrate, very slow generator and generator with predicate that is never satisfied:

```go
package main_test

import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/constraints"
    "github.com/steffnova/go-check/generator"
)

func TestSlowGenerator(t *testing.T) {
    check.Check(t, check.Property(
        func(n int64) error {
            return nil
        },

        // Int64 generator can generate 2^64 values, and likelihood for generated
        // value to satisfy a filter is slim. Generator will keep retrying until
        // Filter's predicate is satisfied. Eventually it will generate it, but it
        // will take a long time to do it, thus making this generator painfully slow.
        generator.Int64().Filter(func(n int64) bool{
            return 0 < n && n < 100
        }),
    ))
}

func TestInfiniteLoopGenerator(t *testing.T) {
    check.Check(t, check.Property(
        func(n int) error {
            return nil
        },

        // Int generator will generate values in range [0, 10], thus never 
        // satisfying Filter's predicate that expectes generated value to
        // be greated than 10. Generator will enter in an endless retrying
        // cycle.
        generator.Int(constraints.Int{
            Min: 0,
            Max: 10,
        }).Filter(func(n int64) bool{
            return n > 10
        }),
    ))
}
```

## Bind

Bind combinator allows binding two generators together using a binder. Binder is a function that accepts one input and has one output. Binder's input must be of the same type as the type of data being generated by it's source and it's output is always `generator.Generator` type. Binder enables using generated data of source generator as an input for another generator. Example below demonstrates how to use Bind to create a generator for 2-element array of integers, where both elements have the same value:

```go
package main_test

import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/constraints"
    "github.com/steffnova/go-check/generator"
)

func TestDuplicate(t *testing.T) {
    // Binder accepts integer n and feeds to array's elements generators.
    binder := func(n int) generator.Generator {
        // ArrayFrom generates array specified by the property ([2]int)
        // Unlike Array, for each element in the array a generator is
        // specified.
        return generator.ArrayFrom(
            // Constat generator always returns a value passed to it.
            generator.Constant(n),
            generator.Constant(n),
        )
    }

    check.Check(t, check.Property(
        func(duplicates [2]int) error {
            if duplicates[0] != duplicates[1] {
                return fmt.Errorf("No duplicates: %v", duplicates)
            }
            return nil
        },
        generator.Int().Bind(binder),
    ))
}
```