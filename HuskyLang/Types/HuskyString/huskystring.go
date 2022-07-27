package HuskyString

import (
	"ofkm.us/husky/HuskyLang/Types"
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
)

func NewString(name string, value string) {
	HuskyProject.CompiledHuskyProject.HuskyStrings = append(HuskyProject.CompiledHuskyProject.HuskyStrings, Types.HuskyString{Name: name, Value: value})
}

func UpdateStringValue(Index int, NewValue string) {
	HuskyProject.CompiledHuskyProject.HuskyStrings[Index].Value = NewValue
}
