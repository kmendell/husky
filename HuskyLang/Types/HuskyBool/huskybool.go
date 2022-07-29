package HuskyBool

import (
	"strings"

	"ofkm.us/husky/HuskyLang/Types"
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
)

func NewBool(Line string) {
	name1 := strings.Split(Line, "(")
	name2 := strings.Split(name1[1], ",")
	val1 := strings.Split(Line, " ")
	val2 := strings.Split(val1[1], ")")
	boolval := false
	if strings.ToLower(val2[0]) == "true" || strings.ToLower(val2[0]) == "yes" || strings.ToLower(val2[0]) == "1" {
		boolval = true
	} else {
		boolval = false
	}
	HuskyProject.CompiledHuskyProject.HuskyBools = append(HuskyProject.CompiledHuskyProject.HuskyBools, Types.HuskyBool{Name: name2[0], Value: boolval})
}

func UpdateBoolValue(Line string) {
	name1 := strings.Split(Line, "(")
	name2 := strings.Split(name1[1], ",")
	val1 := strings.Split(Line, " ")
	val2 := strings.Split(val1[1], ")")
	boolval := false
	if strings.ToLower(val2[0]) == "true" || strings.ToLower(val2[0]) == "yes" || strings.ToLower(val2[0]) == "1" {
		boolval = true
	} else {
		boolval = false
	}
	for _, x := range HuskyProject.CompiledHuskyProject.HuskyBools {
		for j := 0; j < len(HuskyProject.CompiledHuskyProject.HuskyBools); j++ {
			if x.Name == name2[0] {
				HuskyProject.CompiledHuskyProject.HuskyBools[j].Value = boolval
			}
		}
	}

}
