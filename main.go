package main

import (
	"bufio"
	"fmt"
	"os"
)

import "ofkm.us/husky/types"

func huskyOpen() {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for index, line := range fileLines {
		types.HuskyProjectCheck(line, index, filePath)
		types.HuskyProjectStringVaribles(line, index, filePath)
	}

}

func main() {
	huskyOpen()
	// for _, rvar := range compiledHuskyProject.varibles {
	// 	println(rvar.Name, ":", rvar.Value)
	// 	println(rvar.DoesContain("2"))
	// }
}
