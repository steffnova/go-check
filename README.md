# go-check
Property based testing framework for Go. 


The goal is to support generators and shrinkers for all of the go's native types:
  - Numerics
    - Integers: 
        - int8 
        - int16 
        - int32
        - int64 
        - int
    - Unsigned Integers
        - uint8
        - uint16
        - uint32
        - uint64
        - uint
    - Floats
        - float32
        - float64
    - Complex:
        - complex64
        - complex128 
  - Containers
    - Array
    - Slice
    - Map
  - Functions
    - With defined number of input
    - With variadic input
  - Channels
  - Pointers
  - Structures


## Terminology

In order to write any test with go-check 2 things are needed:

``` go
// Property is a function that defines a property by specifing predicate function and generators.
// First parameter predicate must be a function that can have arbitrary number of input parameters
// but can have only one output parameter which must be a GO's error interface. Variadic arbitraries
// parameter defines arbitrary generators used by property to generate random values for each of the 
// predicate's inputs. Number of arbitraries must match number of predicate's inputs
//
// Property will use generators to generate input for predicate function and execute the predicate function.
// If predicate fails property will try to shrink generated input to smalles set of values for which
// predicate function fails. Each type has it's own shrinking strategy.
check.Property(predicate interface{}, ...arbitraries generator.Arbitrary) 

// Check is a function that is used to test the property. By default it will test property 100 times.
check.Check(t *testing.T, property check.property)
```

## Generators

This segment will cover usage of generators for 

### 1.0 Int generators

```go
import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/generator"
    "github.com/steffnova/go-check/constraints"
)

func TestProperty(t *testing.T) {
    // Predicate is a function that accepts two input parameters x, y of type int.
    // Predicate can have only one output parameter and it must error interface.
    // Returning nil indicates predicate success, otherwise failure.
    predicate := func(x, y int) error {
        // Test associativity for integers
        if x + y != y + x {
            return nil
        }
        return fmt.Errorf("associativity failed")
    }

    check.Check(t, check.Property(
        predicate,
        // Int generator with default constraints. Int value generated will be in a
        // range [math.MinInt, math.MaxInt]
        generator.Int(), 
        // Int generator with with custom constraints. Int value generated will be in a
        // range [-5, 5]
        generator.Int(constraints.Int{Min: -5, Max: 5}),
    ))
}
```

### 1.1 Int8, Int16, Int32 and Int64 generators

```go
import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/generator"
    "github.com/steffnova/go-check/constraints"
)

func TestProperty(t *testing.T) {
    predicate := func(x8, x16, x32, x64 int) error {
        return nil
    }

    check.Check(t, check.Property(
        predicate,
        generator.Int8(),  // constraints can be passed: generator.Int8(constraints.Int8{Min: 0, Max: 10})
        generator.Int16(), // constraints can be passed: generator.Int16(constraints.Int16{Min: 0, Max: 10})
        generator.Int32(), // constraints can be passed: generator.Int32(constraints.Int32{Min: 0, Max: 10})
        generator.Int64(), // constraints can be passed: generator.Int64(constraints.Int64{Min: 0, Max: 10})
    ))
}
```

</details> 

### 2.0 First-order functions

With go-check it is possible to define function generator:

```go
import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/generator"
)

func TestProperty(t *testing.T) {
    predicate := func(fn func(x int) int) error {
        if fn(0) != fn (0) {
            return fmt.Errorf("fn is not pure function!")
        }
        return nil
    }

    check.Check(t, check.Property(
        predicate, 
        // Func generator expects generators for output values
        // In this case, generator for Int should be passed to
        // generator.Func vecause fn's type is a function: func(x int) int
        //   
        // Func generator will generate pure function. It will
        // produce same output for same input every time
        generator.Func(generator.Int()),
    ))
}
```

### 2.1 Higher order functions

With go-check it is possible to specify higher order function generators:

```go
import (
    "fmt"
    "testing"

    "github.com/steffnova/go-check"
    "github.com/steffnova/go-check/generator"
)

type higherOrderFn func(int) func(int) int

func TestProperty(t *testing.T) {
    predicate := func(fn higherOrderFn) error {
        if fn(0)(1) != fn (0)(1) {
            return fmt.Errorf("fn is not pure function!")
        }
        return nil
    }

    check.Check(t, check.Property(
        predicate, 
        // Higher order function generator is also a pure function
        generator.Func(generator.Func(generator.Int())),
    ))
}
```