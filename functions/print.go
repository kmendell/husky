package functions

import "ofkm.us/husky/types"

func PrintLogic(varname string) {

	for _, x := range types.CompiledHuskyProject.Varibles {
		if x.Name == varname {
			HuskyPrintLine(x.Name, x.Value)
		} else {
			//TODO :- Just print was is inside quites in print func
		}
	}
}

func HuskyPrintLine(varname string, value string) {
	println(value)
}
