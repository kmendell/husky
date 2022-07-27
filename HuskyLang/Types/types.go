package Types

type HuskyParseObject struct {
	InputFilePath string
	Index         int
	Line          string
	Array         []string
}

type HuskyProject struct {
	Name         string
	HuskyStrings []HuskyString
	HuskyInts    []HuskyInt
	HuskyBools   []HuskyBool
}

type HuskyString struct {
	Name  string
	Value string
}

type HuskyInt struct {
	Name  string
	Value int
}

type HuskyBool struct {
	Name  string
	Value bool
}
