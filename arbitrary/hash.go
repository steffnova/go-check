package arbitrary

import (
	"hash/fnv"
	"reflect"
)

func HashToInt64(vals ...reflect.Value) int64 {
	data := ""
	for _, val := range vals {
		data += EncodeToString(val)
	}

	h64 := fnv.New64()
	h64.Write([]byte(data))
	return int64(h64.Sum64())
}
