package HuskyProject

import "ofkm.us/husky/HuskyType"

type HuskyProject struct {
	Name      string
	Varibles  []HuskyType.HuskyString
	HuskyInts []HuskyType.HuskyInt
}

func NewHuskyProject(dname string, dvar []HuskyType.HuskyString, dint []HuskyType.HuskyInt) HuskyProject {

	project := HuskyProject{Name: dname, Varibles: dvar, HuskyInts: dint}

	return project
}

var CompiledHuskyProject HuskyProject
