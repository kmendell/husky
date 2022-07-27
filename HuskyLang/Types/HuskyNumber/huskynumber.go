package HuskyNumber

import (
	"ofkm.us/husky/HuskyLang/Types"
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
)

func NewNumber(Name string, Value int) {
	HuskyProject.CompiledHuskyProject.HuskyInts = append(HuskyProject.CompiledHuskyProject.HuskyInts, Types.HuskyInt{Name: Name, Value: Value})
}

func UpdateNumberValue(Index int, NewValue int) {
	HuskyProject.CompiledHuskyProject.HuskyInts[Index].Value = NewValue
}
