// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports index/suffixarray'. DO NOT EDIT.

import (
	"index/suffixarray"
	"reflect"
)

func init() {
	Symbols["index/suffixarray"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"New": reflect.ValueOf(suffixarray.New),

		// type definitions
		"Index": reflect.ValueOf((*suffixarray.Index)(nil)),

		// interface wrapper definitions

	}
}
