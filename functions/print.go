package functions

import "ofkm.us/husky/types"

func PrintLogic(varname string) {

	for _, x := range types.CompiledHuskyProject.Varibles {
		if x.Name == varname {
			HuskyPrintLine(x.Name, x.Value)
			break
		} else {
			println("Variable Not Found")
		}
	}
}

func HuskyPrintLine(varname string, value string) {
	println(value)
}