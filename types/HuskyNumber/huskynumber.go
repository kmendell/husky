package HuskyNumber

import "ofkm.us/husky/types"

func NewNumber(Name string, Value int) {
	types.CompiledHuskyProject.HuskyInts = append(types.CompiledHuskyProject.HuskyInts, types.HuskyInt{Name: Name, Value: Value})
}
