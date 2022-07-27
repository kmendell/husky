package HuskyProject

import "ofkm.us/husky/HuskyType"

type HuskyProject struct {
	Name         string
	HuskyStrings []HuskyType.HuskyString
	HuskyInts    []HuskyType.HuskyInt
	HuskyBools   []HuskyType.HuskyBool
}

func NewHuskyProject(dname string, dvar []HuskyType.HuskyString, dint []HuskyType.HuskyInt, dbool []HuskyType.HuskyBool) HuskyProject {

	project := HuskyProject{Name: dname, HuskyStrings: dvar, HuskyInts: dint, HuskyBools: dbool}

	return project
}

var CompiledHuskyProject HuskyProject
