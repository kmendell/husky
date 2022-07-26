package functions

import "ofkm.us/husky/types"

func PrintLogic(varname string) {

	for _, x := range types.CompiledHuskyProject.Varibles {
		if x.Name == varname {
			println(x.Value)
		}
	}

	for _, x := range types.CompiledHuskyProject.HuskyInts {
		if x.Name == varname {
			println(x.Value)
		}
	}
}
