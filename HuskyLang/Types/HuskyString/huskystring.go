package HuskyString

import (
	"strings"

	"ofkm.us/husky/HuskyLang/Types"
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
)

func NewString(Line string) {
	strnamepart1 := strings.Split(Line, "(")
	strnamepart2 := strings.Split(strnamepart1[1], ",")
	strval1 := strings.Split(Line, "\"")
	HuskyProject.CompiledHuskyProject.HuskyStrings = append(HuskyProject.CompiledHuskyProject.HuskyStrings, Types.HuskyString{Name: strnamepart2[0], Value: strval1[1]})
}

func UpdateStringValue(Line string) {
	for _, x := range HuskyProject.CompiledHuskyProject.HuskyStrings {
		strname := strings.Split(Line, "(")
		strname = strings.Split(strname[1], ",")
		updatestr := strings.Split(Line, "\"")
		for j := 0; j < len(HuskyProject.CompiledHuskyProject.HuskyStrings); j++ {
			if x.Name == strname[0] {
				HuskyProject.CompiledHuskyProject.HuskyStrings[j].Value = updatestr[1]
			}
		}
	}
}
