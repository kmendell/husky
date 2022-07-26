package HuskyParse

import (
	"strconv"
	"strings"

	"ofkm.us/husky/HuskyFunctions"
	"ofkm.us/husky/HuskyType"
	"ofkm.us/husky/HuskyType/HuskyNumber"
	"ofkm.us/husky/HuskyType/HuskyStrings"
)

type HuskyParseObject struct {
	InputFilePath string
	Index         int
	Line          string
}

func HuskyParseStart(object HuskyParseObject) {

	if !strings.HasPrefix(object.Line, "//") {
		huskyParseStringVaribles(object.Line, object.Index)
		huskyParseProject(object.Line, object.Index, object.InputFilePath)
		huskyParseIntVaribles(object.Line, object.Index)
		huskyParseFuncPrint(object.Line, object.Index)
	}

}

func huskyParseStringVaribles(line string, index int) {

	if strings.Contains(line, "NewString") {
		strnamepart1 := strings.Split(line, "(")
		strnamepart2 := strings.Split(strnamepart1[1], ",")
		strval1 := strings.Split(line, "\"")
		HuskyStrings.NewString(strnamepart2[0], strval1[1])
	}
}

func huskyParseIntVaribles(line string, index int) {

	if strings.Contains(line, "NewNumber") {
		name1 := strings.Split(line, "(")
		name2 := strings.Split(name1[1], ",")
		val1 := strings.Split(line, " ")
		val2 := strings.Split(val1[1], ")")
		convint, err := strconv.Atoi(val2[0])
		if err != nil {
			panic(err)
		}
		HuskyNumber.NewNumber(name2[0], convint)
	}
}

func huskyParseProject(line string, index int, filePath string) {
	if strings.Contains(line, "project") {
		projName := strings.Split(line, " ")
		if !strings.Contains(projName[1], "main") {
			println("Error in", filePath, "on line", index+1, "\n   Expected Project 'main' found", projName[1])
			// TODO := Create and Throw a new HuskyError
		} else {
			HuskyType.CompiledHuskyProject = HuskyType.NewHuskyProject(projName[1], []HuskyType.HuskyString{}, []HuskyType.HuskyInt{})
		}

	}
}

func huskyParseFuncPrint(line string, index int) {
	if strings.Contains(line, "print") {
		part1 := strings.Split(line, "(")
		part2 := strings.Split(part1[1], ")")
		HuskyFunctions.PrintLogic(part2[0])
	}
}
