package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		huskyProjectCheck(line, index, filePath)
		huskyProjectStringVaribles(line, index, filePath)
	}

}

func main() {
	huskyOpen()
	// for _, rvar := range compiledHuskyProject.varibles {
	// 	println(rvar.Name, ":", rvar.Value)
	// 	println(rvar.DoesContain("2"))
	// }
}
