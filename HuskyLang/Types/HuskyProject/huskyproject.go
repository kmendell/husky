package HuskyProject

import (
	"ofkm.us/husky/HuskyLang/Types"
)

func NewEmptyProject(ProjectName string) Types.HuskyProject {
	project := Types.HuskyProject{Name: ProjectName, HuskyStrings: []Types.HuskyString{}, HuskyInts: []Types.HuskyInt{}, HuskyBools: []Types.HuskyBool{}}

	return project
}

var CompiledHuskyProject Types.HuskyProject
