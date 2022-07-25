package functions

import "ofkm.us/husky/types"

func PrintLogic(varname string) {

	for _, x := range types.CompiledHuskyProject.Varibles {
		if x.Name == varname {
			HuskyPrintLine(x.Name, x.Value)
		} else {
			for _, num := range types.CompiledHuskyProject.HuskyInts {
				if num.Name == varname {
					println(num.Value)
				}
			}
		}
	}
}

func HuskyPrintLine(varname string, value string) {
	println(value)
}
