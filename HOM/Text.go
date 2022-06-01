package HOM

type Text struct {
	ValueLength     int
	SplitTextLength int
	Value           string
	SplitText       []string
}

func (text *Text) CalculateTextHyphens(width, withOffset int) {
	text.SplitText = splitLongText(width-withOffset, text.Value)
	text.computeLengths()
}

func (text *Text) computeLengths() {
	text.ValueLength = len(text.Value)
	text.SplitTextLength = len(text.SplitText)
}
