package HuskyFunctions

import "ofkm.us/husky/HuskyType/HuskyProject"

func PrintLogic(varname string) {

	for _, x := range HuskyProject.CompiledHuskyProject.Varibles {
		if x.Name == varname {
			println(x.Value)
		}
	}

	for _, x := range HuskyProject.CompiledHuskyProject.HuskyInts {
		if x.Name == varname {
			println(x.Value)
		}
	}
}
