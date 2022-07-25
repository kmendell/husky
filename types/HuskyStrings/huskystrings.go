package huskystrings

import "ofkm.us/husky/types"

func NewString(name string, value string) {
	types.CompiledHuskyProject.Varibles = append(types.CompiledHuskyProject.Varibles, types.HuskyString{Name: name, Value: value})
}
