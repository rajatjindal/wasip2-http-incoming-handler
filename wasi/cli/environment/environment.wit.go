// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package environment represents the imported interface "wasi:cli/environment@0.2.0".
package environment

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// GetEnvironment represents the imported function "get-environment".
//
// Get the POSIX-style environment variables.
//
// Each environment variable is provided as a pair of string variable names
// and string value.
//
// Morally, these are a value import, but until value imports are available
// in the component model, this import function should return the same
// values each time it is called.
//
//	get-environment: func() -> list<tuple<string, string>>
//
//go:nosplit
func GetEnvironment() cm.List[[2]string] {
	var result cm.List[[2]string]
	wasmimport_GetEnvironment(&result)
	return result
}

//go:wasmimport wasi:cli/environment@0.2.0 get-environment
//go:noescape
func wasmimport_GetEnvironment(result *cm.List[[2]string])

// GetArguments represents the imported function "get-arguments".
//
// Get the POSIX-style arguments to the program.
//
//	get-arguments: func() -> list<string>
//
//go:nosplit
func GetArguments() cm.List[string] {
	var result cm.List[string]
	wasmimport_GetArguments(&result)
	return result
}

//go:wasmimport wasi:cli/environment@0.2.0 get-arguments
//go:noescape
func wasmimport_GetArguments(result *cm.List[string])

// InitialCWD represents the imported function "initial-cwd".
//
// Return a path that programs should use as their initial current working
// directory, interpreting `.` as shorthand for this.
//
//	initial-cwd: func() -> option<string>
//
//go:nosplit
func InitialCWD() cm.Option[string] {
	var result cm.Option[string]
	wasmimport_InitialCWD(&result)
	return result
}

//go:wasmimport wasi:cli/environment@0.2.0 initial-cwd
//go:noescape
func wasmimport_InitialCWD(result *cm.Option[string])
