package HuskyBool

import (
	"ofkm.us/husky/HuskyLang/Types"
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
)

func NewBool(Name string, Value bool) {
	HuskyProject.CompiledHuskyProject.HuskyBools = append(HuskyProject.CompiledHuskyProject.HuskyBools, Types.HuskyBool{Name: Name, Value: Value})
}

func UpdateBoolValue(Index int, NewValue bool) {
	HuskyProject.CompiledHuskyProject.HuskyBools[Index].Value = NewValue
}
