# go-check
Property based testing framework for Go

## Examples

### 1.0 Numerics

```go
import (
    "fmt"
	"testing"

    "github.com/steffnova/go-check"
	"github.com/steffnova/go-check/generator"
)

func TestProperty(t *testing.T) {
    predicate := func(x, y int) error {
        if x != y {
            return nil
        }
        return fmt.Errorf("x == y")
    }

	check.Check(t, check.Property(
		predicate,
        generator.Int(),
        generator.Int(),
	)
}
```

### 2.0 Strings
```go
import (
    "fmt"
	"testing"

    "github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

func TestProperty(t *testing.T) {
	predicate := func(str string) error {
        result := strings.Join(strings.Split(str, ""))
        if result != str {
            return fmt.Errorf("Splitting and joining a string should return original string")
        }
		return nil
	}

	check.Check(t, check.Property(
        predicate, 
        generator.String(constraints.String{
            // Define Min and Maxi unicode code point that will be used as a range
            // in which values for string runes will be generated
            Rune: constraints.Rune{MinCodePoint: 33, MaxCodePoint: 126},
            // Define Min and Maximum Length of a string
            Length: constraints.Length{Min: 0, Max: 20},
		}),
    ))
}
```

### 3.0 Functions

### 3.1 First-order functions
```go
import (
    "fmt"
	"testing"

    "github.com/steffnova/go-check"
	"github.com/steffnova/go-check/constraints"
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
        // In this case, generator for Int should be passed as
        // fn in predicate is of type: func(int) int
        // Func generator will generate pure function. It will
        // produce same output for same input every time
        generator.Func(generator.Int()),
    )
}
```
