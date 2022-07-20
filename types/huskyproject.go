package types

type HuskyProject struct {
	Name     string
	Varibles []HuskyString
}

func NewHuskyProject(dname string, dvar []HuskyString) HuskyProject {

	project := HuskyProject{Name: dname, Varibles: dvar}

	return project
}

var CompiledHuskyProject HuskyProject

// func HuskyProjectCheck(line string, index int, filePath string) {
// 	if strings.Contains(line, "project") {
// 		projName := strings.Split(line, " ")
// 		if !strings.Contains(projName[1], "main") {
// 			println("Error in", filePath, "on line", index+1, "\n   Expected Project 'main' found", projName[1])
// 		} else {
// 			CompiledHuskyProject = NewHuskyProject(projName[1], []HuskyString{})
// 		}

// 	}
// }

// func HuskyProjectStringVaribles(line string, index int, filePath string) {

// 	if strings.Contains(line, "string") {
// 		sname1 := strings.Split(line, "=")
// 		sname2 := strings.Split(sname1[0], " ")
// 		val1 := strings.Split(sname1[1], "\"")
// 		CompiledHuskyProject.Varibles = append(CompiledHuskyProject.Varibles, HuskyString{Name: sname2[1], Value: val1[1]})
// 	}
// }
