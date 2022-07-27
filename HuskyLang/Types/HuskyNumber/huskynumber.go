package HuskyNumber

import (
	"strconv"
	"strings"

	"ofkm.us/husky/HuskyLang/Types"
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
)

func NewNumber(Line string) {
	name1 := strings.Split(Line, "(")
	name2 := strings.Split(name1[1], ",")
	val1 := strings.Split(Line, " ")
	val2 := strings.Split(val1[1], ")")
	convint, _ := strconv.Atoi(val2[0])
	// HuskyNumber.NewNumber(name2[0], convint)
	HuskyProject.CompiledHuskyProject.HuskyInts = append(HuskyProject.CompiledHuskyProject.HuskyInts, Types.HuskyInt{Name: name2[0], Value: convint})
}

func UpdateNumberValue(Line string) {
	for _, x := range HuskyProject.CompiledHuskyProject.HuskyInts {
		name1 := strings.Split(Line, "(")
		name2 := strings.Split(name1[1], ",")
		val1 := strings.Split(Line, " ")
		val2 := strings.Split(val1[1], ")")
		convint, _ := strconv.Atoi(val2[0])
		for j := 0; j < len(HuskyProject.CompiledHuskyProject.HuskyInts); j++ {
			if x.Name == name2[0] {
				// HuskyNumber.UpdateNumberValue(j, convint)
				HuskyProject.CompiledHuskyProject.HuskyInts[j].Value = convint
			}
		}
	}
}
