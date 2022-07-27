package HuskyProject

import (
	"ofkm.us/husky/HuskyLang/Types"
	"strings"
)

func NewEmptyProject(Line string) {

	projName := strings.Split(Line, " ")
	if !strings.Contains(projName[1], "main") {
		panic(true)
	} else {
		CompiledHuskyProject = Types.HuskyProject{Name: projName[1], HuskyStrings: []Types.HuskyString{}, HuskyInts: []Types.HuskyInt{}, HuskyBools: []Types.HuskyBool{}}
	}

}

var CompiledHuskyProject Types.HuskyProject
