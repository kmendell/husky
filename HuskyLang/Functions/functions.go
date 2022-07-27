package Functions

import (
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
)

func Print(varname string) {

	for _, x := range HuskyProject.CompiledHuskyProject.HuskyStrings {
		if x.Name == varname {
			println(x.Value)
		}
	}

	for _, x := range HuskyProject.CompiledHuskyProject.HuskyInts {
		if x.Name == varname {
			println(x.Value)
		}
	}

	for _, x := range HuskyProject.CompiledHuskyProject.HuskyBools {
		if x.Name == varname {
			println(x.Value)
		}
	}
}
