package Utils

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"ofkm.us/husky/HuskyLang/Parser"
	"ofkm.us/husky/HuskyLang/Types"
)

func OpenFileFromArg() {

	if len(os.Args) != 2 {
		FindHuskyFiles()
	} else {
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
			Parser.HuskyParseFile(Types.HuskyParseObject{InputFilePath: filePath, Index: index, Line: line, Array: fileLines})
		}
	}
}

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		if strings.Contains(s, ".husky") {
			filePath := s
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
				Parser.HuskyParseFile(Types.HuskyParseObject{InputFilePath: filePath, Index: index, Line: line, Array: fileLines})
			}
		}
	}
	return nil
}

func FindHuskyFiles() {
	filepath.WalkDir("./", walk)
}
