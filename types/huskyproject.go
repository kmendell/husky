package types

type HuskyProject struct {
	Name      string
	Varibles  []HuskyString
	HuskyInts []HuskyInt
}

func NewHuskyProject(dname string, dvar []HuskyString, dint []HuskyInt) HuskyProject {

	project := HuskyProject{Name: dname, Varibles: dvar, HuskyInts: dint}

	return project
}

var CompiledHuskyProject HuskyProject
