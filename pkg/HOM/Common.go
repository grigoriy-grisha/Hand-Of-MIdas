package HOM

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func parsePercentStringToPercentFloat(value interface{}) (float64, error) {
	convertedValue, isString := value.(string)

	if !isString {
		return 0, nil
	}

	convertedStringWidth, err := strconv.ParseFloat(strings.TrimSuffix(convertedValue, "%"), 64)

	if err != nil {
		return 0, err
	}

	return convertedStringWidth / 100, nil
}

func assert(t *testing.T, value interface{}, expected interface{}) {
	if value != expected {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf(
			"\033[31m%s:%d: "+"expected result: %d, but get %d"+"\033[39m\n\n",
			append([]interface{}{filepath.Base(file), line}, expected, value)...)
		t.FailNow()
	}
}

func splitLongText(width int, value string) []string {
	var splitText []string

	splitStrings := strings.Split(value, " ")

	preparedString := ""

	for index, splitString := range splitStrings {
		preparedStringLength := len(preparedString + splitString)

		if width <= preparedStringLength {
			splitText = append(splitText, preparedString)
			preparedString = ""
			preparedStringLength = 0
		}

		if preparedStringLength == 0 || index == 0 {
			preparedString += splitString
		} else {
			preparedString += " " + splitString
		}
	}

	if len(preparedString) != 0 {
		splitText = append(splitText, preparedString)
	}

	return splitText
}
