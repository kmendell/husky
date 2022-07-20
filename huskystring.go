package main

import "strings"

type HuskyString struct {
	name  string
	value string
}

func (hs HuskyString) toCaps() string {
	return strings.ToUpper(hs.value)
}

func (hs HuskyString) toLower() string {
	return strings.ToLower(hs.value)
}

func (hs HuskyString) DoesContain(str string) bool {
	return strings.Contains(hs.value, str)
}
