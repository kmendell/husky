package Parser

import (
	"strconv"
	"strings"

	"ofkm.us/husky/HuskyType"
	"ofkm.us/husky/HuskyType/HuskyBool"
	"ofkm.us/husky/HuskyType/HuskyNumber"
	"ofkm.us/husky/HuskyType/HuskyProject"
	"ofkm.us/husky/HuskyType/HuskyStrings"

	"ofkm.us/husky/HuskyLang/Functions"
)

func HuskyParseFile(object HuskyType.HuskyParseObject) {
	line := object.Line

	if !strings.HasPrefix(object.Line, "//") {
		if strings.Contains(line, "NewString") {
			strnamepart1 := strings.Split(line, "(")
			strnamepart2 := strings.Split(strnamepart1[1], ",")
			strval1 := strings.Split(line, "\"")
			HuskyStrings.NewString(strnamepart2[0], strval1[1])
		} else if strings.Contains(line, "UpdateString") {
			for _, x := range HuskyProject.CompiledHuskyProject.HuskyStrings {
				strname := strings.Split(line, "(")
				strname = strings.Split(strname[1], ",")
				updatestr := strings.Split(line, "\"")
				for j := 0; j < len(HuskyProject.CompiledHuskyProject.HuskyStrings); j++ {
					if x.Name == strname[0] {
						HuskyStrings.UpdateStringValue(j, updatestr[1])
					}
				}
			}
		} else if strings.Contains(line, "NewNumber") {
			name1 := strings.Split(line, "(")
			name2 := strings.Split(name1[1], ",")
			val1 := strings.Split(line, " ")
			val2 := strings.Split(val1[1], ")")
			convint, _ := strconv.Atoi(val2[0])
			HuskyNumber.NewNumber(name2[0], convint)
		} else if strings.Contains(line, "UpdateNumber") {
			for _, x := range HuskyProject.CompiledHuskyProject.HuskyInts {
				name1 := strings.Split(line, "(")
				name2 := strings.Split(name1[1], ",")
				val1 := strings.Split(line, " ")
				val2 := strings.Split(val1[1], ")")
				convint, _ := strconv.Atoi(val2[0])
				for j := 0; j < len(HuskyProject.CompiledHuskyProject.HuskyInts); j++ {
					if x.Name == name2[0] {
						HuskyNumber.UpdateNumberValue(j, convint)
					}
				}
			}
		} else if strings.Contains(line, "NewBool") {
			name1 := strings.Split(line, "(")
			name2 := strings.Split(name1[1], ",")
			val1 := strings.Split(line, " ")
			val2 := strings.Split(val1[1], ")")
			boolval := false
			if val2[0] == "true" {
				boolval = true
			} else {
				boolval = false
			}
			HuskyBool.NewBool(name2[0], boolval)
		} else if strings.Contains(line, "UpdateBool") {
			name1 := strings.Split(line, "(")
			name2 := strings.Split(name1[1], ",")
			val1 := strings.Split(line, " ")
			val2 := strings.Split(val1[1], ")")
			boolval := false
			if val2[0] == "true" {
				boolval = true
			} else {
				boolval = false
			}
			for _, x := range HuskyProject.CompiledHuskyProject.HuskyBools {
				for j := 0; j < len(HuskyProject.CompiledHuskyProject.HuskyBools); j++ {
					if x.Name == name2[0] {
						HuskyBool.UpdateBoolValue(j, boolval)
					}
				}
			}

		} else if strings.Contains(line, "project") {
			projName := strings.Split(line, " ")
			if !strings.Contains(projName[1], "main") {
				panic(true)
			} else {
				HuskyProject.CompiledHuskyProject = HuskyProject.NewEmptyProject(projName[1])
			}

		} else if strings.Contains(line, "print") {
			part1 := strings.Split(line, "(")
			part2 := strings.Split(part1[1], ")")
			Functions.Print(part2[0])
		}

	}
}
