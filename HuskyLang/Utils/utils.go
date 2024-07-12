package Utils

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"ofkm.us/husky/HuskyLang/Modules"
	"ofkm.us/husky/HuskyLang/Parser"
	"ofkm.us/husky/HuskyLang/Types"
)

func OpenFileFromArg() {

	if len(os.Args) < 2 {
		FindHuskyFiles()
	} else if len(os.Args) == 3 && os.Args[1] == "-m" {
		EnterHuskyModuleMode()
	} else {
		CheckFirstArgument()
	}
}

func EnterHuskyModuleMode() {
	fmt.Println("Entering Husky Module Mode...")

	if len(os.Args) < 3 {
		fmt.Println("No Module Defined, Please define a module in order to use Husky Module mode.")
	} else {
		if os.Args[2] == "http" {
			Modules.HuskyHTTPModule()
		}
	}

}

func RunHuskyFile(huskyFilePath string) {
	filePath := huskyFilePath
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

func CheckFirstArgument() {
	if os.Args[1] == "-m" {
		EnterHuskyModuleMode()
	} else {
		RunHuskyFile(os.Args[1])
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
