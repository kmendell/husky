package types

func NewHuskyProject(dname string, dvar []HuskyString, dint []HuskyInt) HuskyProject {

	project := HuskyProject{Name: dname, Varibles: dvar, HuskyInts: dint}

	return project
}

var CompiledHuskyProject HuskyProject
