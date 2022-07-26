package HuskyStrings

import "ofkm.us/husky/HuskyType"

import "ofkm.us/husky/HuskyType/HuskyProject"

func NewString(name string, value string) {
	HuskyProject.CompiledHuskyProject.Varibles = append(HuskyProject.CompiledHuskyProject.Varibles, HuskyType.HuskyString{Name: name, Value: value})
}
