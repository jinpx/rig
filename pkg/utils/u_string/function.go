package u_string

import (
	"reflect"
	"runtime"
)

// FunctionName ...
func FunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
