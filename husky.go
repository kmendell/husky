package main

type HuskyProject struct {
	name     string
	varibles []HuskyString
}

func newHuskyProject(dname string, dvar []HuskyString) HuskyProject {

	project := HuskyProject{name: dname, varibles: dvar}

	return project
}

var compiledHuskyProject HuskyProject
