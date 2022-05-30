package HOM

import "strings"

type Text struct {
	ValueLength     int
	SplitTextLength int
	Value           string
	SplitText       []string
}

func (text *Text) CalculateTextHyphens(width, withOffset int) {
	text.SplitText = text.splitLongText(width-withOffset, text.Value)
	text.computeLengths()
}

func (text *Text) computeLengths() {
	text.ValueLength = len(text.Value)
	text.SplitTextLength = len(text.SplitText)
}

//todo возможно, можно вынести в отдельную структуру
func (text *Text) splitLongText(width int, value string) []string {
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
