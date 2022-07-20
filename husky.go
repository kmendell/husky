package main

import "ofkm.us/husky/types"

type HuskyProject struct {
	name     string
	varibles []types.HuskyString
}

func newHuskyProject(dname string, dvar []types.HuskyString) HuskyProject {

	project := HuskyProject{name: dname, varibles: dvar}

	return project
}

var compiledHuskyProject HuskyProject
