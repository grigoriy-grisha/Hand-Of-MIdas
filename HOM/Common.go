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

func assert(t *testing.T, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		t.FailNow()
	}
}
