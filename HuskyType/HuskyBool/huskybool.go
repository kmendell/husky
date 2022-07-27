package HuskyBool

import "ofkm.us/husky/HuskyType"

import "ofkm.us/husky/HuskyType/HuskyProject"

func NewBool(Name string, Value bool) {
	HuskyProject.CompiledHuskyProject.HuskyBools = append(HuskyProject.CompiledHuskyProject.HuskyBools, HuskyType.HuskyBool{Name: Name, Value: Value})
}

func UpdateBoolValue(Index int, NewValue bool) {
	HuskyProject.CompiledHuskyProject.HuskyBools[Index].Value = NewValue
}
