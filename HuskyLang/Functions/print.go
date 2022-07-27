package Functions

import (
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
	"strings"
)

func PrintParse(Line string) {
	part1 := strings.Split(Line, "(")
	part2 := strings.Split(part1[1], ")")
	Println(part2[0])
}

func Println(varname string) {

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
