package HuskyStrings

import "ofkm.us/husky/HuskyType"

func NewString(name string, value string) {
	HuskyType.CompiledHuskyProject.Varibles = append(HuskyType.CompiledHuskyProject.Varibles, HuskyType.HuskyString{Name: name, Value: value})
}
