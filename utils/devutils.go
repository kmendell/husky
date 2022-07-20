package utils

import "ofkm.us/husky/types"

func DevforLoop() {
	for _, rvar := range types.CompiledHuskyProject.Varibles {
		println(rvar.Name, ":", rvar.Value)
		println(rvar.DoesContain("2"))
	}
}
