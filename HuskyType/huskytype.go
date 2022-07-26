package HuskyType

type HuskyInt struct {
	Name  string
	Value int
}

type HuskyString struct {
	Name  string
	Value string
}

type HuskyProject struct {
	Name      string
	Varibles  []HuskyString
	HuskyInts []HuskyInt
}
