package utils

import (
	"strconv"
	"strings"

	"ofkm.us/husky/functions"
	"ofkm.us/husky/types"
	"ofkm.us/husky/types/HuskyNumber"
	"ofkm.us/husky/types/HuskyStrings"
)

type HuskyParseObject struct {
	Type          string
	InputFilePath string
	Index         int
	Line          string
}

func HuskyParseStart(object HuskyParseObject) {

	if !strings.HasPrefix(object.Line, "//") {
		switch object.Type {
		case types.GetType().HuskyString:
			huskyParseStringVaribles(object.Line, object.Index)
		case types.GetType().HuskyProject:
			huskyParseProject(object.Line, object.Index, object.InputFilePath)
		case types.GetType().HuskyInt:
			huskyParseIntVaribles(object.Line, object.Index)
		case types.GetType().HuskyFuncPrint:
			huskyParseFuncPrint(object.Line, object.Index)
		default:
			break
		}
	}

}

func huskyParseStringVaribles(line string, index int) {
	if strings.Contains(line, "string.new()") {
		sname1 := strings.Split(line, "=")
		sname2 := strings.Split(sname1[0], " ")
		val1 := strings.Split(sname1[1], "\"")
		types.CompiledHuskyProject.Varibles = append(types.CompiledHuskyProject.Varibles, types.HuskyString{Name: sname2[1], Value: val1[1]})
	}

	if strings.Contains(line, "NewString") {
		strnamepart1 := strings.Split(line, "(")
		strnamepart2 := strings.Split(strnamepart1[1], ",")
		strval1 := strings.Split(line, "\"")
		HuskyStrings.NewString(strnamepart2[0], strval1[1])
	}
}

func huskyParseIntVaribles(line string, index int) {
	if strings.Contains(line, "int.new()") {
		sname1 := strings.Split(line, "=")
		sname2 := strings.Split(sname1[0], " ")
		val1 := strings.Split(sname1[1], " ")
		dint, err := strconv.Atoi(val1[1])
		if err != nil {
			panic(err)
		}
		types.CompiledHuskyProject.HuskyInts = append(types.CompiledHuskyProject.HuskyInts, types.HuskyInt{Name: sname2[1], Value: dint})
	}

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
			types.CompiledHuskyProject = types.NewHuskyProject(projName[1], []types.HuskyString{}, []types.HuskyInt{})
		}

	}
}

func huskyParseFuncPrint(line string, index int) {
	if strings.Contains(line, "print") {
		part1 := strings.Split(line, "(")
		part2 := strings.Split(part1[1], ")")
		functions.PrintLogic(part2[0])
	}
}
