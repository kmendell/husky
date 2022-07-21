package utils

// These functions are really only to be used while devloping Husky, upon release this will be taken out of the compiled binary

import "ofkm.us/husky/types"

func DevforLoop() {
	for _, rvar := range types.CompiledHuskyProject.HuskyInts {
		println(rvar.Name, ":", rvar.Value)
	}
}
