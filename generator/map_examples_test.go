package generator_test

import (
	"fmt"

	"github.com/steffnova/go-check/constraints"
	"github.com/steffnova/go-check/generator"
)

// This example demonstrates usage of Map(Int8(), Bool()), generator for generation of map[int8]bool
// values. Map() generator requires generators for generating map's key-value pairs. In this example
// Int8() is used as a key generator and bool is used as value generator.
func ExampleMap() {
	streamer := generator.Streamer(
		func(m map[int8]bool) {
			fmt.Printf("%#v\n", m)
		},
		generator.Map(
			generator.Int8(),
			generator.Bool(),
		),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// map[int8]bool{-127:false, -120:true, -118:true, -116:false, -102:true, -81:false, -76:false, -69:true, -68:true, -57:false, -50:false, -40:false, -30:true, -25:true, -16:true, -10:false, -4:true, 0:false, 7:false, 11:false, 19:true, 21:false, 33:true, 36:true, 40:false, 48:false, 72:false, 79:false, 99:false, 101:false, 126:false}
	// map[int8]bool{-117:true, -94:false, -73:false, -42:false, -21:false, -12:false, 14:true, 27:false, 28:false, 35:false, 38:false, 39:true, 43:false, 45:false, 52:true, 75:false, 116:true}
	// map[int8]bool{-127:true, -126:true, -125:false, -123:false, -120:true, -119:false, -115:true, -112:false, -111:false, -109:true, -106:true, -103:true, -101:true, -95:false, -90:true, -88:true, -85:false, -83:false, -81:false, -79:false, -78:false, -76:true, -72:false, -70:true, -69:false, -68:false, -63:false, -58:true, -55:false, -54:true, -51:true, -47:true, -45:true, -43:false, -42:false, -40:false, -36:false, -33:true, -29:false, -28:false, -25:true, -23:false, -19:false, -13:true, -12:true, -10:true, -3:true, -1:true, 1:false, 3:true, 8:true, 12:true, 14:true, 15:false, 20:false, 26:false, 28:false, 30:false, 31:false, 32:false, 34:true, 39:false, 40:true, 41:true, 43:false, 46:true, 47:true, 50:false, 56:true, 62:true, 69:true, 71:true, 76:false, 77:true, 78:false, 79:false, 84:false, 92:true, 95:true, 98:true, 99:true, 100:true, 106:false, 108:false, 109:true, 111:true, 113:true, 114:true, 118:true, 119:true, 122:false, 123:true, 127:false}
	// map[int8]bool{-108:false, -97:false, -92:false, -90:false, -85:false, -82:true, -68:false, -61:true, -56:true, -40:true, -27:false, -11:true, 12:false, 20:true, 30:true, 49:false, 56:true, 61:false, 70:true, 78:false, 80:true, 82:false, 103:false, 104:false, 111:false, 119:true}
	// map[int8]bool{-128:true, -126:false, -124:false, -122:true, -120:false, -119:false, -116:false, -113:true, -112:true, -109:false, -107:true, -104:true, -100:true, -98:false, -91:false, -88:false, -82:true, -81:false, -80:false, -79:false, -76:false, -68:false, -67:false, -63:false, -61:false, -57:true, -53:true, -52:false, -49:true, -44:false, -42:true, -41:false, -29:false, -24:false, -21:false, -18:true, -13:false, -12:true, -10:false, -9:true, -8:true, -6:true, -3:false, -2:true, 0:true, 1:true, 3:true, 8:true, 12:true, 16:false, 17:false, 18:false, 21:true, 24:false, 28:false, 34:true, 36:false, 38:true, 42:true, 46:false, 51:false, 52:true, 59:true, 68:false, 69:false, 71:false, 72:false, 76:false, 77:true, 82:true, 85:true, 88:false, 89:true, 90:true, 91:false, 94:true, 95:false, 96:true, 100:false, 101:false, 106:true, 108:true, 111:false, 115:false, 120:false, 122:false, 123:false, 125:false}
	// map[int8]bool{-128:true, -104:false, -96:true, -85:false, -69:true, -68:false, -54:true, -42:false, -27:true, 0:true, 20:true, 29:true, 44:true, 93:false, 114:true, 126:true}
	// map[int8]bool{-128:true, -57:false, -40:true, -6:true, 13:false, 38:false}
	// map[int8]bool{-110:true, -102:false, -90:true, -75:false, -58:true, -57:false, -55:true, -52:false, -49:false, -1:false, 2:true, 17:true, 33:false, 35:true, 45:true, 47:true, 51:false, 56:true, 66:false, 67:true, 68:true, 70:true, 72:false, 86:true, 91:true, 105:false, 115:false, 118:true, 120:false, 123:false}
	// map[int8]bool{-124:false, -123:true, -121:true, -120:true, -115:true, -111:true, -106:true, -103:false, -99:true, -98:false, -94:false, -91:true, -77:false, -74:false, -71:false, -69:false, -68:true, -66:false, -61:false, -59:false, -58:true, -56:false, -48:false, -41:true, -36:false, -35:true, -34:true, -30:true, -24:false, -21:true, -19:false, -16:false, -14:false, -12:false, -9:false, -6:true, 0:false, 4:false, 8:true, 11:false, 12:false, 17:true, 20:false, 21:true, 22:false, 35:true, 37:true, 40:true, 43:false, 44:false, 54:false, 56:true, 58:true, 60:false, 61:false, 62:false, 72:true, 83:false, 86:false, 91:false, 92:false, 96:true, 111:true, 112:true, 115:true, 117:true, 120:true, 123:false}
	// map[int8]bool{-97:true, -75:false, -69:false, -46:true, -44:true, -8:true, 19:false, 27:true, 35:false, 36:true, 61:true, 87:true, 106:false, 107:true, 115:false, 123:false, 124:true}
}

// This example demonstrates usage of Map(Int8(), Uint8())generator with constraints for generation
// of map[int8]uint8 values. Map() generator requires generators for generating map's key-value pairs.
// Int8() and Uint8() are used as map's key and value generators (respectively). Constraints define
// range of generatable values for map's size.
func ExampleMap_constraints() {
	streamer := generator.Streamer(
		func(m map[int8]uint8) {
			fmt.Printf("%#v\n", m)
		},
		generator.Map(
			generator.Int8(),
			generator.Uint8(),
			constraints.Length{
				Min: 0,
				Max: 5,
			},
		),
	)

	if err := generator.Stream(0, 10, streamer); err != nil {
		panic(err)
	}
	// Output:
	// map[int8]uint8{-40:0xa9, -31:0x2d, 21:0xcc, 40:0x68, 79:0xdf}
	// map[int8]uint8{}
	// map[int8]uint8{-67:0xef}
	// map[int8]uint8{-127:0xed, -102:0xa2, -57:0x40, 19:0xed}
	// map[int8]uint8{-120:0xad, -10:0x1b, 72:0xe0}
	// map[int8]uint8{}
	// map[int8]uint8{-54:0xab, -19:0x26, 4:0x73, 126:0xd7}
	// map[int8]uint8{-25:0xaa}
	// map[int8]uint8{-76:0x2e}
	// map[int8]uint8{-127:0x69, -10:0x71, 7:0x98, 18:0x9c, 101:0x28}
}
