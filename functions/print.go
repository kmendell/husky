package functions

import "ofkm.us/husky/types"

func PrintLogic(varname string) {

	for _, x := range types.CompiledHuskyProject.Varibles {
		if x.Name == varname {
			HuskyPrintLine(x.Name, x.Value)
		} else {
			break
		}
	}

	for _, x := range types.CompiledHuskyProject.HuskyInts {
		if x.Name == varname {
			println(x.Value)
		} else {
			break
		}
	}
}

func HuskyPrintLine(varname string, value string) {
	println(value)
}
