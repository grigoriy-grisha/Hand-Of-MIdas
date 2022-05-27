package consoleRenderer

import (
	"awesomeProject/HOM"
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

func (tr *textRenderer) renderText() {
	bounding := tr.element.Bounding
	SplitText := tr.element.Text.SplitText
	TextValue := tr.element.Text.Value

	if tr.isLeftTopAlign() {
		for textIndex, spitText := range SplitText {
			y := bounding.OffsetTopLeft.Y + 1 + textIndex

			if y >= bounding.OffsetBottomLeft.Y {
				break
			}

			for i, textItem := range spitText {
				x := bounding.OffsetTopLeft.X + 1 + i

				termbox.SetCell(x, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isCenterLeftAlign() {
		centerPositionY := tr.computeCenterY(len(SplitText) - 1)

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

	if tr.isLeftBottomAlign() {
		length := len(SplitText)

		for textIndex, spitText := range SplitText {
			y := bounding.OffsetBottomLeft.Y - length + textIndex

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			for i, textItem := range spitText {
				x := bounding.OffsetTopLeft.X + 1 + i

				termbox.SetCell(x, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isRightTopAlign() {
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

	if tr.isCenterRightAlign() {
		centerPositionY := tr.computeCenterY(len(SplitText))

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

	if tr.isRightBottomAlign() {
		length := len(SplitText)

		for textIndex, spitText := range SplitText {

			y := bounding.OffsetBottomRight.Y - length + textIndex

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

	if tr.isCenterTopAlign() {
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

	if tr.isCenterBottomAlign() {
		length := len(SplitText)

		for textIndex, splitText := range SplitText {
			y := bounding.OffsetBottomRight.Y - length + textIndex

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			centerByX := tr.computeCenterX(len(splitText) - 1)

			for i, textItem := range splitText {
				termbox.SetCell(centerByX+i, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isCenterCenterAlign() {
		centerPositionY := tr.computeCenterY(len(SplitText) - 1)

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
}

func TextIsNotEmpty(text string) bool {
	return len(text) != 0
}

// todo заменить назвение center на middle чтобы не путаться
func (tr *textRenderer) isLeftTopAlign() bool {
	return tr.alignContent == HOM.AlignContentLeft && tr.verticalContent == HOM.VerticalContentTop
}

func (tr *textRenderer) isCenterLeftAlign() bool {
	return tr.alignContent == HOM.AlignContentLeft && tr.verticalContent == HOM.VerticalContentCenter
}

func (tr *textRenderer) isRightTopAlign() bool {
	return tr.alignContent == HOM.AlignContentRight && tr.verticalContent == HOM.VerticalContentTop
}

func (tr *textRenderer) isLeftBottomAlign() bool {
	return tr.alignContent == HOM.AlignContentLeft && tr.verticalContent == HOM.VerticalContentBottom
}

func (tr *textRenderer) isCenterRightAlign() bool {
	return tr.alignContent == HOM.AlignContentRight && tr.verticalContent == HOM.VerticalContentCenter
}

func (tr *textRenderer) isRightBottomAlign() bool {
	return tr.alignContent == HOM.AlignContentRight && tr.verticalContent == HOM.VerticalContentBottom
}

func (tr *textRenderer) isCenterTopAlign() bool {
	return tr.alignContent == HOM.AlignContentCenter && tr.verticalContent == HOM.VerticalContentTop
}

func (tr *textRenderer) isCenterBottomAlign() bool {
	return tr.alignContent == HOM.AlignContentCenter && tr.verticalContent == HOM.VerticalContentBottom
}

func (tr *textRenderer) isCenterCenterAlign() bool {
	return tr.alignContent == HOM.AlignContentCenter && tr.verticalContent == HOM.VerticalContentCenter
}

func (tr *textRenderer) computeCenterX(textLength int) int {
	return (tr.element.Bounding.ClientTopLeft.X + (tr.element.Style.Width / 2)) - textLength/2
}

func (tr *textRenderer) computeCenterY(rows int) int {
	return (tr.element.Bounding.OffsetTopLeft.Y+tr.element.Bounding.OffsetBottomRight.Y)/2 - rows/2
}
