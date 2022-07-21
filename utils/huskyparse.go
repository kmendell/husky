package utils

import (
	"strconv"
	"strings"

	"ofkm.us/husky/types"
)

type HuskyParseObject struct {
	Type          string
	InputFilePath string
	Index         int
	Line          string
}

func HuskyParseStart(object HuskyParseObject) {

	switch object.Type {
	case types.GetType().HuskyString:
		huskyParseStringVaribles(object.Line, object.Index)
	case types.GetType().HuskyProject:
		huskyParseProject(object.Line, object.Index, object.InputFilePath)
	case types.GetType().HuskyInt:
		huskyParseIntVaribles(object.Line, object.Index)
	default:
		break
	}

}

func huskyParseStringVaribles(line string, index int) {
	if strings.Contains(line, "string") {
		sname1 := strings.Split(line, "=")
		sname2 := strings.Split(sname1[0], " ")
		val1 := strings.Split(sname1[1], "\"")
		types.CompiledHuskyProject.Varibles = append(types.CompiledHuskyProject.Varibles, types.HuskyString{Name: sname2[1], Value: val1[1]})
	}
}

func huskyParseIntVaribles(line string, index int) {
	if strings.Contains(line, "int") {
		sname1 := strings.Split(line, "=")
		sname2 := strings.Split(sname1[0], " ")
		val1 := strings.Split(sname1[1], " ")
		dint, err := strconv.Atoi(val1[1])
		if err != nil {
			panic(err)
		}
		types.CompiledHuskyProject.HuskyInts = append(types.CompiledHuskyProject.HuskyInts, types.HuskyInt{Name: sname2[1], Value: dint})
	}
}

func huskyParseProject(line string, index int, filePath string) {
	if strings.Contains(line, "project") {
		projName := strings.Split(line, " ")
		if !strings.Contains(projName[1], "main") {
			println("Error in", filePath, "on line", index+1, "\n   Expected Project 'main' found", projName[1])
		} else {
			types.CompiledHuskyProject = types.NewHuskyProject(projName[1], []types.HuskyString{}, []types.HuskyInt{})
		}

	}
}
