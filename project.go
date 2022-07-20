package main

import "strings"

func huskyProjectCheck(line string, index int, filePath string) {
	if strings.Contains(line, "project") {
		projName := strings.Split(line, " ")
		if !strings.Contains(projName[1], "main") {
			println("Error in", filePath, "on line", index+1, "\n   Expected Project 'main' found", projName[1])
		} else {
			compiledHuskyProject = newHuskyProject(projName[1], []HuskyString{})
		}

	}
}

func huskyProjectStringVaribles(line string, index int, filePath string) {

	if strings.Contains(line, "string") {
		sname1 := strings.Split(line, "=")
		sname2 := strings.Split(sname1[0], " ")
		val1 := strings.Split(sname1[1], "\"")
		compiledHuskyProject.varibles = append(compiledHuskyProject.varibles, HuskyString{name: sname2[1], value: val1[1]})
	}
}
