package types

type HuskyInt struct {
	Name  string
	Value int
}

type HuskyString struct {
	Name  string
	Value string
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
