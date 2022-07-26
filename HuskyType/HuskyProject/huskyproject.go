package HuskyProject

import "ofkm.us/husky/HuskyType"

type HuskyProject struct {
	Name         string
	HuskyStrings []HuskyType.HuskyString
	HuskyInts    []HuskyType.HuskyInt
}

func NewHuskyProject(dname string, dvar []HuskyType.HuskyString, dint []HuskyType.HuskyInt) HuskyProject {

	project := HuskyProject{Name: dname, HuskyStrings: dvar, HuskyInts: dint}

	return project
}

var CompiledHuskyProject HuskyProject
