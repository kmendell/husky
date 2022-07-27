package main

import (
	"bufio"
	"fmt"
	"os"

	"ofkm.us/husky/HuskyParse"
)

func huskyOpenFile() {
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
		HuskyParse.HuskyParseStart(HuskyParse.HuskyParseObject{InputFilePath: filePath, Index: index, Line: line, Array: fileLines})
	}

}

func main() {
	huskyOpenFile()
}
