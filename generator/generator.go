package generator

// // arbitrary.Arbitrary returns new arbitrary and it's coresponding shrinker
// type arbitrary.Arbitrary func() (arbitrary.Arbitrary, shrinker.Shrinker)

// // arbitrary.Generator returns arbitrary.Arbitrary for a type specified by "target" parameter, that can be used to
// // generate target's value using "bias" and "r".
// type arbitrary.Generator func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error)

// // Map (combinator) returns generator that maps generated value to a new one using mapper. Mapper
// // must be a function that has one input and one output. Mapper's input type must match generated
// // value's type and Mapper's output type must match generator's target type. Error is returned if
// // mapper is invalid or if generator of mapper's input type returns an error.
// func (generator arbitrary.Generator) Map(mapper interface{}) arbitrary.Generator {
// 	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
// 		val := reflect.ValueOf(mapper)
// 		switch {
// 		case val.Kind() != reflect.Func:
// 			return nil, fmt.Errorf("%w. Mapper must be a function", ErrorMapper)
// 		case val.Type().NumOut() != 1:
// 			return nil, fmt.Errorf("%w. Mapper must have 1 output value", ErrorMapper)
// 		case val.Type().NumIn() != 1:
// 			return nil, fmt.Errorf("%w. Mapper must have 1 input value", ErrorMapper)
// 		case val.Type().Out(0).Kind() != target.Kind():
// 			return nil, fmt.Errorf("%w. Mappers output kind: %s must match target's kind. Got: %s", ErrorMapper, val.Type().Out(0).Kind(), target.Kind())
// 		}

// 		generate, err := generator(val.Type().In(0), bias, r)
// 		if err != nil {
// 			return nil, fmt.Errorf("Failed to use base generator. %w", err)
// 		}

// 		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
// 			arb, shrinker := generate()
// 			return arbitrary.Arbitrary{
// 				Value:      reflect.ValueOf(mapper).Call([]reflect.Value{arb.Value})[0],
// 				Precursors: []arbitrary.Arbitrary{arb},
// 			}, shrinker.Map(mapper)
// 		}, nil
// 	}
// }

// // Filter (combinator) returns a generator that generates value only if predicate is satisfied.
// // Predicate is a function that has one input and one output. Predicate's input parameter must
// // match generator's target type, and output parameter must be bool. Error is returned
// // predicate is invalid, or generator for predicate's input returns an error.
// //
// // NOTE: Returned generator will retry generation until predicate is satisfied, which
// // can impact the generator's speed.
// func (generator arbitrary.Generator) Filter(predicate interface{}) arbitrary.Generator {
// 	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
// 		generate, err := generator(target, bias, r)
// 		switch val := reflect.ValueOf(predicate); {
// 		case err != nil:
// 			return nil, fmt.Errorf("Failed to use base generator. %w", err)
// 		case val.Kind() != reflect.Func:
// 			return nil, fmt.Errorf("%w. Filter predicate must be a function", ErrorFilter)
// 		case val.Type().NumIn() != 1:
// 			return nil, fmt.Errorf("%w. Filter predicate must have one input value", ErrorFilter)
// 		case val.Type().NumOut() != 1:
// 			return nil, fmt.Errorf("%w. Filter predicate must have one output value", ErrorFilter)
// 		case val.Type().Out(0).Kind() != reflect.Bool:
// 			return nil, fmt.Errorf("%w. Filter predicate must have bool as a output type", ErrorFilter)
// 		}

// 		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {
// 			for {
// 				arb, shrinker := generate()
// 				outputs := reflect.ValueOf(predicate).Call([]reflect.Value{arb.Value})
// 				if outputs[0].Bool() {
// 					return arb, shrinker.Filter(predicate)
// 				}
// 			}
// 		}, nil
// 	}
// }

// // Bind (combinator) returns bounded generator using "binder" parameter. Binder
// // is a function that has one input and one output. Input's type must match
// // generated value's type. Output type must be generator.arbitrary.Generator. Binder allows
// // using generated value of one generator as an input to another generator. Error
// // is returned if binder is invalid, generator returns an error or bound generator
// // returns an error.
// func (generator arbitrary.Generator) Bind(binder interface{}) arbitrary.Generator {
// 	return func(target reflect.Type, bias constraints.Bias, r arbitrary.Random) (arbitrary.Arbitrary, error) {
// 		binderVal := reflect.ValueOf(binder)
// 		switch t := reflect.TypeOf(binder); {
// 		case t.Kind() != reflect.Func:
// 			return nil, fmt.Errorf("%w. Binder must be a function", ErrorBinder)
// 		case t.NumIn() != 1:
// 			return nil, fmt.Errorf("%w. Binder must have one input value", ErrorBinder)
// 		case t.NumOut() != 1:
// 			return nil, fmt.Errorf("%w. Binder must have one output values", ErrorBinder)
// 		case t.Out(0) != reflect.TypeOf(arbitrary.Generator(nil)):
// 			return nil, fmt.Errorf("%w. Binder's output type must be generator.arbitrary.Generator", ErrorBinder)
// 		}

// 		generate, err := generator(binderVal.Type().In(0), bias, r)
// 		if err != nil {
// 			return nil, fmt.Errorf("Failed to use base generator: %w", err)
// 		}

// 		binder := func(source arbitrary.Arbitrary) (arbitrary.Arbitrary, shrinker.Shrinker, error) {
// 			generator := binderVal.Call([]reflect.Value{source.Value})[0].Interface().(arbitrary.Generator)
// 			generate, err := generator(target, bias, r)
// 			if err != nil {
// 				return arbitrary.Arbitrary{}, nil, fmt.Errorf("arbitrary.Generator Binding failed (%s -> %s): %w", binderVal.Type().In(0), target, err)
// 			}

// 			val, shrinker := generate()
// 			val.Precursors = append(val.Precursors, source)
// 			return val, shrinker, nil
// 		}

// 		sourceArb, sourceShrinker := generate()
// 		boundVal, boundShrinker, err := binder(sourceArb)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return func() (arbitrary.Arbitrary, shrinker.Shrinker) {

// 			return boundVal, sourceShrinker.
// 				Bind(binder, boundVal, boundShrinker, boundShrinker)
// 		}, nil
// 	}
// }
