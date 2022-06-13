package consoleRenderer

import (
	"awesomeProject/pkg/HOM"
	"github.com/nsf/termbox-go"
)

type textRenderer struct {
	alignContent    HOM.AlignContent
	verticalContent HOM.VerticalContent
	element         *HOM.Element
}

type NewTextRendererParams struct {
	alignContent    HOM.AlignContent
	verticalContent HOM.VerticalContent
	element         *HOM.Element
}

func NewTextRenderer(params NewTextRendererParams) *textRenderer {
	return &textRenderer{
		alignContent:    params.alignContent,
		verticalContent: params.verticalContent,
		element:         params.element,
	}
}

//todo termbox.SetCell можно внести
//todo сделать text рендерер независимым от termbox
//todo рефакторинг

func printText(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

//todo высчитать все это зарание

func (tr *textRenderer) renderLeftAlignText() {
	SplitText := tr.element.Text.SplitText
	bounding := tr.element.Bounding
	SplitTextLength := tr.element.Text.SplitTextLength

	if isTopVerticalContent(tr.verticalContent) {
		for textIndex, spitText := range SplitText {
			y := bounding.OffsetTopLeft.Y + textIndex + 1

			if y >= bounding.OffsetBottomLeft.Y {
				break
			}

			printText(bounding.OffsetTopLeft.X+1, y, termbox.ColorWhite, termbox.ColorBlack, spitText)
		}
	}

	if isCenterVerticalContent(tr.verticalContent) {
		centerPositionY := tr.computeCenterY(SplitTextLength - 1)

		for textIndex, spitText := range SplitText {
			y := centerPositionY + textIndex

			if y >= bounding.OffsetBottomLeft.Y {
				continue
			}

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			for i, textItem := range spitText {
				x := bounding.OffsetBottomLeft.X + 1 + i
				termbox.SetCell(x, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if isBottomVerticalContent(tr.verticalContent) {

		for textIndex, spitText := range SplitText {
			y := bounding.OffsetBottomLeft.Y - SplitTextLength + textIndex

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			for i, textItem := range spitText {
				x := bounding.OffsetTopLeft.X + 1 + i

				termbox.SetCell(x, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}
}

func (tr *textRenderer) renderRightAlignText() {
	SplitText := tr.element.Text.SplitText
	bounding := tr.element.Bounding
	TextValue := tr.element.Text.Value
	SplitTextLength := tr.element.Text.SplitTextLength

	if isTopVerticalContent(tr.verticalContent) {
		for textIndex, spitText := range SplitText {
			textLength := len(spitText) - 1

			y := bounding.OffsetTopLeft.Y + 1 + textIndex

			if y >= bounding.OffsetBottomRight.Y {
				break
			}

			for i := textLength; i >= 0; i-- {

				x := bounding.OffsetBottomRight.X - 1 - i

				termbox.SetCell(x, y, rune(TextValue[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if isCenterVerticalContent(tr.verticalContent) {
		centerPositionY := tr.computeCenterY(SplitTextLength)

		for textIndex, spitText := range SplitText {
			y := centerPositionY + textIndex

			if y >= bounding.OffsetBottomRight.Y {
				continue
			}

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			textLength := len(spitText) - 1

			for i := textLength; i >= 0; i-- {
				x := bounding.OffsetBottomRight.X - 1 - i
				termbox.SetCell(x, y, rune(TextValue[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if isBottomVerticalContent(tr.verticalContent) {
		for textIndex, spitText := range SplitText {

			y := bounding.OffsetBottomRight.Y - SplitTextLength + textIndex

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			textLength := len(spitText) - 1

			for i := textLength; i >= 0; i-- {
				x := bounding.OffsetBottomRight.X - 1 - i
				termbox.SetCell(x, y, rune(TextValue[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}
}

func (tr *textRenderer) renderCenterAlignText() {
	SplitText := tr.element.Text.SplitText
	bounding := tr.element.Bounding
	SplitTextLength := tr.element.Text.SplitTextLength

	if isTopVerticalContent(tr.verticalContent) {
		for textIndex, splitText := range SplitText {
			y := bounding.OffsetTopLeft.Y + 1 + textIndex

			if y >= bounding.OffsetBottomRight.Y {
				break
			}

			centerByX := tr.computeCenterX(len(splitText) - 1)

			for i, textItem := range splitText {
				termbox.SetCell(centerByX+i, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}

	}

	if isCenterVerticalContent(tr.verticalContent) {
		centerPositionY := tr.computeCenterY(SplitTextLength - 1)

		for textIndex, spitText := range SplitText {
			y := centerPositionY + textIndex

			if y >= bounding.OffsetBottomRight.Y {
				continue
			}

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			startPosition := tr.computeCenterX(len(spitText) - 1)

			for i, textItem := range spitText {
				termbox.SetCell(startPosition+i, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if isBottomVerticalContent(tr.verticalContent) {

		for textIndex, splitText := range SplitText {
			y := bounding.OffsetBottomRight.Y - SplitTextLength + textIndex

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			centerByX := tr.computeCenterX(len(splitText) - 1)

			for i, textItem := range splitText {
				termbox.SetCell(centerByX+i, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

}

func (tr *textRenderer) renderText() {
	if isLeftAlignContent(tr.alignContent) {
		tr.renderLeftAlignText()
		return
	}

	if isRightAlignContent(tr.alignContent) {
		tr.renderRightAlignText()
		return
	}

	if isCenterAlignContent(tr.alignContent) {
		tr.renderCenterAlignText()
		return
	}

}

func TextIsNotEmpty(text *HOM.Text) bool {
	if text == nil {
		return false
	}
	return len(text.Value) != 0
}

func isLeftAlignContent(alignContent HOM.AlignContent) bool {
	return alignContent == HOM.AlignContentLeft
}

func isRightAlignContent(alignContent HOM.AlignContent) bool {
	return alignContent == HOM.AlignContentRight
}

func isCenterAlignContent(alignContent HOM.AlignContent) bool {
	return alignContent == HOM.AlignContentCenter
}

func isTopVerticalContent(verticalContent HOM.VerticalContent) bool {
	return verticalContent == HOM.VerticalContentTop
}

func isCenterVerticalContent(verticalContent HOM.VerticalContent) bool {
	return verticalContent == HOM.VerticalContentCenter
}

func isBottomVerticalContent(verticalContent HOM.VerticalContent) bool {
	return verticalContent == HOM.VerticalContentBottom
}

func (tr *textRenderer) computeCenterX(textLength int) int {
	return (tr.element.Bounding.OffsetTopLeft.X + (tr.element.Bounding.Width / 2)) - textLength/2
}

func (tr *textRenderer) computeCenterY(rows int) int {
	return (tr.element.Bounding.OffsetTopLeft.Y+tr.element.Bounding.OffsetBottomRight.Y)/2 - rows/2
}
