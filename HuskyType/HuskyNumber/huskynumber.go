package HuskyNumber

import "ofkm.us/husky/HuskyType"

import "ofkm.us/husky/HuskyType/HuskyProject"

func NewNumber(Name string, Value int) {
	HuskyProject.CompiledHuskyProject.HuskyInts = append(HuskyProject.CompiledHuskyProject.HuskyInts, HuskyType.HuskyInt{Name: Name, Value: Value})
}
