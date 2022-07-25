package types

type HuskyInt struct {
	Name  string
	Value int
}

type HuskyString struct {
	Name  string
	Value string
}

type HuskyError struct {
	ErrorCode   int
	StringValue string
}

type HuskyProject struct {
	Name      string
	Varibles  []HuskyString
	HuskyInts []HuskyInt
}

type HuskyType struct {
	HuskyString    string
	HuskyInt       string
	HuskyProject   string
	HuskyError     string
	HuskyFuncPrint string
}

func GetType() HuskyType {

	htype := HuskyType{
		HuskyString:    "HuskyString",
		HuskyInt:       "HuskyInt",
		HuskyProject:   "HuskyProject",
		HuskyError:     "HuskyError",
		HuskyFuncPrint: "HuskyFuncPrint",
	}

	return htype
}

func NewHuskyProject(dname string, dvar []HuskyString, dint []HuskyInt) HuskyProject {

	project := HuskyProject{Name: dname, Varibles: dvar, HuskyInts: dint}

	return project
}

var CompiledHuskyProject HuskyProject
