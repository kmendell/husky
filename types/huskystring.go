package types

import "strings"

type HuskyString struct {
	Name  string
	Value string
}

func (hs HuskyString) toCaps() string {
	return strings.ToUpper(hs.Value)
}

func (hs HuskyString) toLower() string {
	return strings.ToLower(hs.Value)
}

func (hs HuskyString) DoesContain(str string) bool {
	return strings.Contains(hs.Value, str)
}
