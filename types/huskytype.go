package types

type HuskyType struct {
	HuskyString  string
	HuskyInt     string
	HuskyProject string
	HuskyError   string
}

func GetType() HuskyType {
	htype := HuskyType{HuskyString: "HuskyString", HuskyInt: "HuskyInt", HuskyProject: "HuskyProject", HuskyError: "HuskyError"}

	return htype
}
