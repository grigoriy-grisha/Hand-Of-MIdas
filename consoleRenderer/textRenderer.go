package consoleRenderer

import (
	"awesomeProject/HOM"
	"github.com/nsf/termbox-go"
)

type textRenderer struct {
	alignContent     HOM.AlignContent
	verticalContent  HOM.VerticalContent
	topX             int
	topY             int
	height           int
	width            int
	botY             int
	botX             int
	normalizedHeight int
	normalizeTopX    int
}

type NewTextRendererParams struct {
	alignContent    HOM.AlignContent
	verticalContent HOM.VerticalContent
	topX            int
	topY            int
	height          int
	width           int
	paddingBottom   int
	paddingTop      int
	paddingLeft     int
	paddingRight    int
}

func NewTextRenderer(params NewTextRendererParams) *textRenderer {
	return &textRenderer{
		alignContent:     params.alignContent,
		verticalContent:  params.verticalContent,
		height:           params.height,
		normalizedHeight: params.height - params.paddingBottom,
		normalizeTopX:    params.topX + params.paddingLeft,
		topX:             params.topX,
		topY:             params.topY + params.paddingTop,
		botY:             params.topY + params.height - params.paddingBottom,
		//может rightX
		botX:  params.topX + params.width - params.paddingRight,
		width: params.width,
	}
}

//todo termbox.SetCell можно внести
//todo сделать text рендерер независимым от termbox
//todo рефакторинг

func (tr *textRenderer) renderText(element HOM.Element) {
	//normalizedHeight := element.Style.Height - element.Style.PaddingTop - element.Style.PaddingBottom

	if tr.isLeftTopAlign() {
		for textIndex, spitText := range element.Text.SplitText {
			y := tr.topY + 1 + textIndex

			if y >= tr.botY {
				break
			}

			for i, textItem := range spitText {
				x := tr.normalizeTopX + 1 + i

				termbox.SetCell(x, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isCenterLeftAlign() {
		centerPositionY := tr.computeCenterY(len(element.Text.SplitText) - 1)

		for textIndex, spitText := range element.Text.SplitText {
			y := centerPositionY + textIndex

			if y >= tr.botY {
				continue
			}

			if y <= tr.topY {
				continue
			}

			for i, textItem := range spitText {
				x := tr.normalizeTopX + 1 + i
				termbox.SetCell(x, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isLeftBottomAlign() {
		length := len(element.Text.SplitText)

		for textIndex, spitText := range element.Text.SplitText {
			y := tr.botY - length + textIndex

			if y <= tr.topY {
				continue
			}

			for i, textItem := range spitText {
				x := tr.normalizeTopX + 1 + i

				termbox.SetCell(x, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isRightTopAlign() {
		for textIndex, spitText := range element.Text.SplitText {
			textLength := len(spitText) - 1

			y := tr.topY + 1 + textIndex

			if y >= tr.normalizedHeight {
				break
			}

			for i := textLength; i >= 0; i-- {

				x := tr.botX - 1 - i

				termbox.SetCell(x, y, rune(element.Text.Value[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isCenterRightAlign() {
		centerPositionY := tr.computeCenterY(len(element.Text.SplitText))

		for textIndex, spitText := range element.Text.SplitText {
			y := centerPositionY + textIndex

			if y >= tr.botY {
				continue
			}

			if y <= tr.topY {
				continue
			}

			textLength := len(spitText) - 1

			for i := textLength; i >= 0; i-- {
				x := tr.botX - 1 - i
				termbox.SetCell(x, y, rune(element.Text.Value[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isRightBottomAlign() {
		length := len(element.Text.SplitText)

		for textIndex, spitText := range element.Text.SplitText {

			y := tr.botY - length + textIndex

			if y <= tr.topY {
				continue
			}

			textLength := len(spitText) - 1

			for i := textLength; i >= 0; i-- {
				x := tr.botX - 1 - i
				termbox.SetCell(x, y, rune(element.Text.Value[textLength-i]), termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isCenterTopAlign() {
		for textIndex, splitText := range element.Text.SplitText {
			y := tr.topY + 1 + textIndex

			if y >= tr.normalizedHeight {
				break
			}

			centerByX := tr.computeCenterX(len(splitText) - 1)

			for i, textItem := range splitText {
				termbox.SetCell(centerByX+i, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}

	}

	if tr.isCenterBottomAlign() {
		length := len(element.Text.SplitText)

		for textIndex, splitText := range element.Text.SplitText {
			y := tr.botY - length + textIndex

			if y <= tr.topY {
				continue
			}

			centerByX := tr.computeCenterX(len(splitText) - 1)

			for i, textItem := range splitText {
				termbox.SetCell(centerByX+i, y, textItem, termbox.ColorWhite, termbox.ColorBlack)
			}
		}
	}

	if tr.isCenterCenterAlign() {
		centerPositionY := tr.computeCenterY(len(element.Text.SplitText) - 1)

		for textIndex, spitText := range element.Text.SplitText {
			y := centerPositionY + textIndex

			if y >= tr.botY {
				continue
			}

			if y <= tr.topY {
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
	return (tr.topX + (tr.width / 2)) - textLength/2
}

func (tr *textRenderer) computeCenterY(rows int) int {
	return (tr.topY+tr.botY)/2 - rows/2
}

//func (textRenderer *textRenderer) handleSplitTextLeftSide(x int, y int, topLeftY int, element HOM.Element) {
//	for textIndex, spitText := range element.Text.SplitText {
//		if y+textIndex >= topLeftY+element.Style.Height {
//			break
//		}
//
//		for i, textItem := range spitText {
//
//			termbox.SetCell(x+i, y+textIndex, textItem, termbox.ColorWhite, termbox.ColorBlack)
//		}
//	}
//}
