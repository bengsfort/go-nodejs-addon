package main

/*
#cgo CFLAGS: -I./vendor
#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
#cgo linux LDFLAGS: -Wl,--unresolved-symbols=ignore-all
#cgo windows LDFLAGS: -Wl,--allow-shlib-undefined
#include <node_api.h>
*/
import "C"

//export napi_register_module_v1
func napi_register_module_v1(env C.napi_env, exports C.napi_value) C.napi_value {
	var world_str C.napi_value

	C.napi_create_string_utf8(env, C.CString("World"), C.size_t(5), &world_str)
	C.napi_set_named_property(env, exports, C.CString("hello"), world_str)

	return exports
}

func main() {}
