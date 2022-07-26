package HuskyFunctions

import "ofkm.us/husky/HuskyType"

func PrintLogic(varname string) {

	for _, x := range HuskyType.CompiledHuskyProject.Varibles {
		if x.Name == varname {
			println(x.Value)
		}
	}

	for _, x := range HuskyType.CompiledHuskyProject.HuskyInts {
		if x.Name == varname {
			println(x.Value)
		}
	}
}
