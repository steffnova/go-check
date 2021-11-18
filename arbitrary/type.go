package arbitrary

import "reflect"

var Type = reflect.TypeOf([]interface{}{}).Elem()
