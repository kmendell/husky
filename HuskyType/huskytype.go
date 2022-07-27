package HuskyType

type HuskyInt struct {
	Name  string
	Value int
}

type HuskyString struct {
	Name  string
	Value string
}

type HuskyBool struct {
	Name  string
	Value bool
}

type HuskyParseObject struct {
	InputFilePath string
	Index         int
	Line          string
	Array         []string
}
