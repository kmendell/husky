package HuskyNumber

import "ofkm.us/husky/HuskyType"

func NewNumber(Name string, Value int) {
	HuskyType.CompiledHuskyProject.HuskyInts = append(HuskyType.CompiledHuskyProject.HuskyInts, HuskyType.HuskyInt{Name: Name, Value: Value})
}
