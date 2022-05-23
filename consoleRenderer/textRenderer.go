package consoleRenderer

import (
	"awesomeProject/Dom"
	"github.com/nsf/termbox-go"
	"strings"
)

type TextRenderer struct {
	alignContent    Dom.AlignContent
	verticalContent Dom.VerticalContent
}

//todo termbox.SetCell можно внести
//todo сделать text рендерер независимым от termbox

func (textRenderer *TextRenderer) renderText(element Dom.Element) {
	topLeftX := element.Style.X
	topLeftY := element.Style.Y

	bottomLeftY := topLeftY + element.Style.Height
	bottomRightX := element.Style.X + element.Style.Width

	if textRenderer.isLeftTopAlign() {
		textRenderer.handleSplitTextLeftSide(topLeftX+1, topLeftY+1, topLeftY, element)
	}

	if textRenderer.isCenterLeftAlign() {
		//todo внести в отдельную функцию
		splitText := textRenderer.splitLongText(element.Style.Width, element.Text)
		centerPositionY := (topLeftY+bottomLeftY)/2 - len(splitText)/2

		for textIndex, spitText := range splitText {
			if centerPositionY+textIndex >= topLeftY+element.Style.Height {
				break
			}

			for i, textItem := range spitText {
				termbox.SetCell(topLeftX+1+i, centerPositionY+textIndex, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if textRenderer.isLeftBottomAlign() {
		for textIndex, spitText := range textRenderer.splitLongText(element.Style.Width, element.Text) {
			if bottomLeftY-1+textIndex >= topLeftY+element.Style.Height {
				break
			}
			for i, textItem := range spitText {
				termbox.SetCell(topLeftX+1+i, bottomLeftY-1+textIndex, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if textRenderer.isRightTopAlign() {
		for textIndex, spitText := range textRenderer.splitLongText(element.Style.Width, element.Text) {
			if topLeftY+1+textIndex >= topLeftY+element.Style.Height {
				break
			}

			textLength := len(spitText) - 1

			for i := textLength; i >= 0; i-- {
				termbox.SetCell(bottomRightX-1-i, topLeftY+1+textIndex, rune(element.Text[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if textRenderer.isCenterRightAlign() {
		splitText := textRenderer.splitLongText(element.Style.Width, element.Text)
		splitTextLength := len(splitText)

		for textIndex, spitText := range splitText {
			centerPositionY := (topLeftY+bottomLeftY)/2 - splitTextLength/2

			if centerPositionY+textIndex >= topLeftY+element.Style.Height {
				break
			}

			textLength := len(spitText) - 1

			for i := textLength; i >= 0; i-- {
				termbox.SetCell(bottomRightX-1-i, centerPositionY+textIndex, rune(element.Text[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if textRenderer.isRightBottomAlign() {
		for textIndex, spitText := range textRenderer.splitLongText(element.Style.Width, element.Text) {
			if bottomLeftY-1+textIndex >= topLeftY+element.Style.Height {
				break
			}

			textLength := len(spitText) - 1

			for i := textLength; i >= 0; i-- {
				termbox.SetCell(bottomRightX-1-i, bottomLeftY-1+textIndex, rune(element.Text[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if textRenderer.isCenterTopAlign() {

		for textIndex, spitText := range textRenderer.splitLongText(element.Style.Width, element.Text) {
			if topLeftY+1+textIndex >= topLeftY+element.Style.Height {
				break
			}

			textLength := len(spitText) - 1

			startPosition := (element.Style.X + (element.Style.Width / 2)) - textLength/2

			for i, textItem := range spitText {
				termbox.SetCell(startPosition+i, topLeftY+1+textIndex, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}

	}

	if textRenderer.isCenterBottomAlign() {
		for textIndex, spitText := range textRenderer.splitLongText(element.Style.Width, element.Text) {
			if bottomLeftY-1+textIndex >= topLeftY+element.Style.Height {
				break
			}

			textLength := len(spitText) - 1

			startPosition := (element.Style.X + (element.Style.Width / 2)) - textLength/2

			for i, textItem := range spitText {
				termbox.SetCell(startPosition+i, bottomLeftY-1, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if textRenderer.isCenterCenterAlign() {
		splitText := textRenderer.splitLongText(element.Style.Width, element.Text)
		centerPositionY := (topLeftY+bottomLeftY)/2 - len(splitText)/2

		for textIndex, spitText := range splitText {
			if centerPositionY+textIndex >= topLeftY+element.Style.Height {
				break
			}

			textLength := len(spitText) - 1

			startPosition := (element.Style.X + (element.Style.Width / 2)) - textLength/2

			for i, textItem := range spitText {
				termbox.SetCell(startPosition+i, centerPositionY+textIndex, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}

	}
}

func TextIsNotEmpty(text string) bool {
	return len(text) != 0
}

// todo заменить назвение center на middle чтобы не путаться
func (textRenderer *TextRenderer) isLeftTopAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentLeft && textRenderer.verticalContent == Dom.VerticalContentTop
}

func (textRenderer *TextRenderer) isCenterLeftAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentLeft && textRenderer.verticalContent == Dom.VerticalContentCenter
}

func (textRenderer *TextRenderer) isRightTopAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentRight && textRenderer.verticalContent == Dom.VerticalContentTop
}

func (textRenderer *TextRenderer) isLeftBottomAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentLeft && textRenderer.verticalContent == Dom.VerticalContentBottom
}

func (textRenderer *TextRenderer) isCenterRightAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentRight && textRenderer.verticalContent == Dom.VerticalContentCenter
}

func (textRenderer *TextRenderer) isRightBottomAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentRight && textRenderer.verticalContent == Dom.VerticalContentBottom
}

func (textRenderer *TextRenderer) isCenterTopAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentCenter && textRenderer.verticalContent == Dom.VerticalContentTop
}

func (textRenderer *TextRenderer) isCenterBottomAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentCenter && textRenderer.verticalContent == Dom.VerticalContentBottom
}

func (textRenderer *TextRenderer) isCenterCenterAlign() bool {
	return textRenderer.alignContent == Dom.AlignContentCenter && textRenderer.verticalContent == Dom.VerticalContentCenter
}

// todo разделить по строкам текст, рефткоринг

func (textRenderer *TextRenderer) splitLongText(width int, text string) []string {
	var splitText []string

	splitStrings := strings.Split(text, " ")

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

func (textRenderer *TextRenderer) handleSplitTextLeftSide(x int, y int, topLeftY int, element Dom.Element) {
	for textIndex, spitText := range textRenderer.splitLongText(element.Style.Width, element.Text) {
		if y+textIndex >= topLeftY+element.Style.Height {
			break
		}

		for i, textItem := range spitText {

			termbox.SetCell(x+i, y+textIndex, textItem, termbox.ColorWhite, termbox.ColorBlack)
		}
	}
}
