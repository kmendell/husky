package HuskyStrings

import "ofkm.us/husky/HuskyType"

import "ofkm.us/husky/HuskyType/HuskyProject"

func NewString(name string, value string) {
	HuskyProject.CompiledHuskyProject.HuskyStrings = append(HuskyProject.CompiledHuskyProject.HuskyStrings, HuskyType.HuskyString{Name: name, Value: value})
}

func UpdateStringValue(Index int, NewValue string) {
	HuskyProject.CompiledHuskyProject.HuskyStrings[Index].Value = NewValue
}
