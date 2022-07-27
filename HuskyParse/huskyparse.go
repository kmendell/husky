package HuskyParse

import (
	"strconv"
	"strings"

	"ofkm.us/husky/HuskyFunctions"
	"ofkm.us/husky/HuskyType"
	"ofkm.us/husky/HuskyType/HuskyBool"
	"ofkm.us/husky/HuskyType/HuskyNumber"
	"ofkm.us/husky/HuskyType/HuskyProject"
	"ofkm.us/husky/HuskyType/HuskyStrings"
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
				strname2 := strings.Split(strname[1], ",")
				updatestr := strings.Split(line, "\"")
				for j := 0; j < len(HuskyProject.CompiledHuskyProject.HuskyStrings); j++ {
					if x.Name == strname2[0] {
						HuskyStrings.UpdateStringValue(j, updatestr[1])
					}
				}
			}
		} else if strings.Contains(line, "NewNumber") {
			name1 := strings.Split(line, "(")
			name2 := strings.Split(name1[1], ",")
			val1 := strings.Split(line, " ")
			val2 := strings.Split(val1[1], ")")
			convint, err := strconv.Atoi(val2[0])
			if err != nil {
				panic(err)
			}
			HuskyNumber.NewNumber(name2[0], convint)
		} else if strings.Contains(line, "UpdateNumber") {
			for _, x := range HuskyProject.CompiledHuskyProject.HuskyInts {
				name1 := strings.Split(line, "(")
				name2 := strings.Split(name1[1], ",")
				val1 := strings.Split(line, " ")
				val2 := strings.Split(val1[1], ")")
				convint, err := strconv.Atoi(val2[0])
				if err != nil {
					panic(err)
				}
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
		} else if strings.Contains(line, "project") {
			projName := strings.Split(line, " ")
			if !strings.Contains(projName[1], "main") {
				println("Error in", object.InputFilePath, "on line", object.Index+1, "\n   Expected Project 'main' found", projName[1])
				panic(true)
				// TODO := Create and Throw a new HuskyError
			} else {
				HuskyProject.CompiledHuskyProject = HuskyProject.NewHuskyProject(projName[1], []HuskyType.HuskyString{}, []HuskyType.HuskyInt{}, []HuskyType.HuskyBool{})
			}

		} else if strings.Contains(line, "print") {
			part1 := strings.Split(line, "(")
			part2 := strings.Split(part1[1], ")")
			HuskyFunctions.Print(part2[0])
		}

	}
}
