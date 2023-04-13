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
	// map[int8]bool{-128:false, -116:true, -100:false, -91:true, -86:true, -76:false, -68:false, -64:true, -62:true, -57:false, -54:true, -51:true, -40:false, -38:true, -21:true, -11:false, 0:true, 4:false, 7:false, 10:false, 19:true, 23:true, 25:false, 30:true, 67:false, 73:false, 76:false, 120:false, 122:true, 126:true, 127:false}
	// map[int8]bool{-119:false, -90:false, -84:true, -83:true, -74:false, -67:false, -53:false, -48:false, -45:false, -42:false, -41:true, -33:false, -28:true, -16:false, 12:false, 18:true, 23:true, 35:false, 36:false, 40:true, 46:true, 61:false, 66:true, 67:false, 68:false, 73:true, 75:false, 81:true, 94:false, 97:true, 99:false, 106:false, 109:true, 115:true, 123:true, 127:false}
	// map[int8]bool{-128:false, -126:true, -118:true, -117:false, -106:true, -105:false, -99:false, -87:true, -84:true, -78:true, -77:true, -76:false, -71:false, -68:false, -64:false, -63:false, -58:true, -55:false, -53:true, -47:false, -43:false, -40:false, -37:false, -35:false, -31:false, -30:true, -26:false, -22:false, -14:true, -13:false, -3:false, -1:false, 0:true, 1:false, 12:true, 15:false, 23:false, 25:true, 26:true, 28:false, 29:true, 33:true, 34:true, 39:true, 45:false, 47:true, 50:false, 53:true, 54:false, 60:false, 71:false, 72:true, 77:false, 84:true, 90:true, 98:true, 101:true, 104:true, 109:true, 111:true, 113:false, 115:false, 117:false, 119:false, 123:false, 125:false, 127:false}
	// map[int8]bool{-126:true, -123:true, -122:false, -120:true, -119:true, -118:true, -112:true, -110:true, -108:false, -100:true, -95:true, -88:true, -86:true, -84:false, -73:true, -69:true, -68:false, -62:true, -58:true, -56:false, -46:false, -45:true, -40:false, -39:true, -36:false, -35:true, -34:false, -33:true, -30:true, -26:true, -21:false, -18:true, -15:true, -11:true, -1:false, 0:false, 8:true, 12:true, 13:true, 17:true, 24:false, 25:false, 30:false, 31:false, 38:true, 39:false, 41:true, 43:true, 47:false, 52:true, 58:true, 62:true, 68:true, 69:false, 74:false, 75:true, 76:true, 79:false, 81:false, 89:true, 92:false, 94:false, 95:true, 98:true, 101:true, 103:true, 106:false, 108:true, 111:false, 115:true, 118:true}
	// map[int8]bool{-128:false, -127:true, -123:true, -122:false, -121:false, -116:false, -114:false, -110:false, -107:true, -106:true, -104:false, -99:true, -91:true, -87:false, -85:false, -84:false, -81:false, -61:false, -56:true, -55:false, -53:false, -51:true, -50:true, -49:true, -45:true, -44:true, -41:true, -36:false, -35:false, -34:true, -32:false, -30:true, -26:true, -23:false, -21:true, -20:true, -18:true, -12:true, -5:true, 0:true, 3:false, 8:true, 11:true, 12:true, 13:false, 14:true, 16:false, 26:true, 31:false, 35:true, 41:true, 42:true, 44:false, 47:true, 51:false, 52:true, 55:true, 61:true, 64:false, 68:true, 70:true, 75:true, 78:false, 80:true, 81:false, 82:true, 83:true, 86:true, 88:true, 89:false, 90:false, 93:false, 94:false, 95:false, 97:false, 98:true, 103:false, 108:false, 114:false, 119:true, 120:true, 121:true, 126:true}
	// map[int8]bool{-128:false, -127:false, -125:true, -120:false, -118:true, -116:false, -115:true, -114:false, -110:false, -109:true, -108:false, -107:true, -106:true, -100:true, -99:false, -95:false, -94:false, -93:false, -92:false, -88:false, -86:false, -84:true, -80:true, -79:false, -77:false, -75:true, -72:true, -69:true, -67:false, -66:false, -54:false, -53:true, -52:false, -49:true, -48:true, -47:false, -46:true, -45:false, -44:true, -42:false, -40:false, -39:true, -37:false, -36:false, -35:true, -32:true, -30:false, -28:false, -24:false, -21:true, -20:false, -19:true, -18:false, -16:true, -15:false, -12:true, -10:true, -8:true, -6:false, -5:false, -3:true, 0:false, 1:false, 2:true, 3:false, 8:true, 12:true, 20:false, 25:true, 27:true, 28:false, 31:true, 36:false, 38:false, 44:true, 46:false, 47:false, 48:true, 49:true, 50:true, 54:false, 57:false, 60:true, 63:true, 67:true, 68:false, 76:false, 77:false, 80:true, 88:false, 96:false, 100:true, 102:true, 104:false, 105:true, 118:true, 119:false, 122:true, 124:false}
	// map[int8]bool{-123:false, -121:true, -115:false, -113:false, -112:false, -111:false, -104:false, -95:true, -82:false, -79:false, -74:false, -73:true, -72:true, -69:true, -65:true, -60:false, -59:false, -58:true, -57:true, -54:true, -51:true, -40:true, -33:true, -25:true, -22:false, -21:false, -19:false, -14:false, -12:true, -3:false, 1:false, 2:true, 6:false, 8:false, 12:false, 14:true, 16:false, 18:true, 19:true, 31:true, 35:false, 41:false, 44:false, 48:false, 50:true, 53:false, 61:true, 63:true, 66:true, 67:true, 68:true, 72:false, 76:true, 84:true, 98:true, 107:true, 110:false, 119:false, 120:false, 122:false}
	// map[int8]bool{-128:false, -127:false, -126:false, -124:false, -123:true, -120:true, -117:false, -115:false, -114:false, -112:true, -110:false, -106:false, -105:false, -100:true, -97:false, -95:false, -94:false, -90:false, -83:false, -80:true, -77:false, -75:false, -74:false, -71:false, -66:false, -64:false, -62:true, -61:false, -55:false, -52:true, -50:true, -44:true, -42:false, -41:true, -40:true, -38:false, -36:false, -35:false, -34:false, -33:false, -22:false, -21:true, -19:true, -18:true, -14:false, -11:false, -8:true, -7:true, -4:true, -3:true, -1:false, 5:true, 8:true, 9:true, 11:false, 12:true, 17:false, 19:true, 21:true, 24:false, 28:false, 30:false, 35:true, 37:true, 40:true, 41:false, 42:true, 44:true, 49:false, 54:true, 58:false, 66:false, 68:false, 69:false, 72:true, 76:true, 80:true, 81:false, 82:true, 85:true, 89:true, 92:true, 95:false, 96:false, 97:false, 98:false, 108:false, 111:false, 112:true, 113:true, 117:false, 121:false, 122:true, 123:true}
	// map[int8]bool{-128:false, -125:true, -117:true, -112:true, -110:false, -106:true, -102:false, -90:false, -81:false, -79:true, -72:false, -71:true, -70:false, -69:false, -67:true, -66:false, -65:true, -58:true, -47:false, -46:true, -36:false, -34:false, -28:false, -26:true, -23:true, -16:false, -14:false, -8:true, -7:false, -6:true, -3:true, 0:true, 1:false, 4:false, 8:false, 14:true, 19:false, 27:false, 29:true, 32:false, 40:false, 43:true, 49:false, 50:true, 51:true, 52:false, 61:true, 62:true, 63:false, 69:false, 70:true, 73:false, 77:true, 82:true, 86:true, 90:true, 92:false, 93:false, 95:true, 99:false, 103:true, 104:false, 106:false, 112:false, 113:true, 119:true, 121:true, 124:false, 125:false, 127:true}
	// map[int8]bool{-127:true, -125:true, -123:true, -122:false, -119:false, -115:false, -103:true, -102:false, -99:false, -96:false, -86:true, -85:false, -81:true, -79:false, -77:true, -75:true, -70:false, -69:true, -66:false, -62:true, -57:false, -53:true, -52:false, -50:true, -49:false, -39:true, -24:true, -21:false, -19:true, -17:false, -15:false, -13:true, -11:false, -7:false, -5:true, -3:false, -1:true, 1:true, 4:false, 5:false, 6:false, 12:true, 18:true, 21:true, 22:false, 26:false, 28:false, 38:true, 40:false, 42:true, 44:false, 49:true, 51:false, 53:true, 57:true, 59:false, 66:false, 72:true, 73:true, 85:false, 90:true, 99:false, 101:true, 102:false, 103:true, 110:true, 111:false, 114:false, 116:false, 118:false, 119:true, 122:true}
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
	// map[int8]uint8{-92:0x15, -81:0xfb, -16:0x1e, 40:0xbe, 122:0x4f}
	// map[int8]uint8{40:0x49, 67:0x24}
	// map[int8]uint8{-102:0xcd, -57:0x1e, 23:0x5d}
	// map[int8]uint8{-76:0x8b, 120:0x13, 127:0x56}
	// map[int8]uint8{-64:0x1b, -54:0x71, -36:0xaa, -4:0x84, 10:0x50}
	// map[int8]uint8{-126:0x91, 19:0xbd, 30:0x24}
	// map[int8]uint8{25:0xb8}
	// map[int8]uint8{-128:0xfe, -62:0x51, 76:0x16}
	// map[int8]uint8{}
	// map[int8]uint8{-116:0x71, 36:0xc5}
}
