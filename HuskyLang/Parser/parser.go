package Parser

import (
	"strings"

	"ofkm.us/husky/HuskyLang/Functions"

	"ofkm.us/husky/HuskyLang/Types"
	"ofkm.us/husky/HuskyLang/Types/HuskyBool"
	"ofkm.us/husky/HuskyLang/Types/HuskyNumber"
	"ofkm.us/husky/HuskyLang/Types/HuskyProject"
	"ofkm.us/husky/HuskyLang/Types/HuskyString"
)

func HuskyParseFile(object Types.HuskyParseObject) {
	line := object.Line

	if !strings.HasPrefix(object.Line, "//") {
		if strings.Contains(line, "project") {
			HuskyProject.NewEmptyProject(line)
		} else if strings.Contains(line, "NewString") {
			HuskyString.NewString(line)
		} else if strings.Contains(line, "UpdateString") {
			HuskyString.UpdateStringValue(line)
		} else if strings.Contains(line, "NewNumber") {
			HuskyNumber.NewNumber(line)
		} else if strings.Contains(line, "UpdateNumber") {
			HuskyNumber.UpdateNumberValue(line)
		} else if strings.Contains(line, "NewBool") {
			HuskyBool.NewBool(line)
		} else if strings.Contains(line, "UpdateBool") {
			HuskyBool.UpdateBoolValue(line)
		} else if strings.Contains(line, "print") {
			Functions.PrintParse(line)
		}

	}
}
