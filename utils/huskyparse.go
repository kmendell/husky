package utils

import "strings"

import "ofkm.us/husky/types"

type HuskyParseObject struct {
	Type          string
	InputFilePath string
	Index         int
	Line          string
}

func HuskyParseStart(object HuskyParseObject) {
	if strings.Contains(object.Type, "HuskyString") {
		huskyParseStringVaribles(object.Line, object.Index)
	} else {
		println("Not a String")
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

func huksyParseIntVaribles() {

}

func huskyParseProject() {

}
