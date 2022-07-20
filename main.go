package main

import (
	"bufio"
	"fmt"
	"os"
)

import "ofkm.us/husky/types"
import "ofkm.us/husky/utils"

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
	utils.DevforLoop()
}
