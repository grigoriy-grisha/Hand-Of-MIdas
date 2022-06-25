package HOMR

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

func (tr *textRenderer) renderLeftAlignText() {
	bounding := tr.element.Bounding
	SplitTextLength := tr.element.Text.SplitTextLength

	if isTopVerticalContent(tr.verticalContent) {
		PrintTextStart(
			PrintTextParams{
				drawX:       bounding.OffsetTopLeft.X + 1,
				drawY:       bounding.OffsetTopLeft.Y,
				boundStartY: bounding.OffsetTopLeft.Y,
				boundEndY:   bounding.OffsetBottomLeft.Y,
				fg:          termbox.ColorWhite,
				bg:          termbox.ColorBlack,
				splitText:   tr.element.Text.SplitText,
			},
		)
	}

	if isCenterVerticalContent(tr.verticalContent) {
		PrintTextStart(
			PrintTextParams{
				drawX:         bounding.OffsetBottomLeft.X + 1,
				drawY:         tr.computeCenterY(SplitTextLength) - 1,
				boundStartY:   bounding.OffsetTopLeft.Y,
				boundEndY:     bounding.OffsetBottomLeft.Y,
				fg:            termbox.ColorWhite,
				bg:            termbox.ColorBlack,
				splitText:     tr.element.Text.SplitText,
				textDirection: 0,
			},
		)
	}

	if isBottomVerticalContent(tr.verticalContent) {
		PrintTextStart(
			PrintTextParams{
				drawX:         bounding.OffsetBottomLeft.X + 1,
				drawY:         bounding.OffsetBottomLeft.Y - SplitTextLength - 1,
				boundEndY:     bounding.OffsetBottomLeft.Y,
				boundStartY:   bounding.OffsetTopLeft.Y,
				fg:            termbox.ColorWhite,
				bg:            termbox.ColorBlack,
				splitText:     tr.element.Text.SplitText,
				textDirection: 0,
			},
		)
	}
}

func (tr *textRenderer) renderRightAlignText() {
	bounding := tr.element.Bounding
	SplitTextLength := tr.element.Text.SplitTextLength

	if isTopVerticalContent(tr.verticalContent) {
		PrintTextEnd(
			PrintTextParams{
				drawX:         bounding.OffsetBottomRight.X + 1,
				drawY:         bounding.OffsetTopLeft.Y,
				boundEndY:     bounding.OffsetBottomRight.Y,
				boundStartY:   bounding.OffsetTopLeft.Y,
				fg:            termbox.ColorWhite,
				bg:            termbox.ColorBlack,
				splitText:     tr.element.Text.SplitText,
				textDirection: 1,
			},
		)
	}

	if isCenterVerticalContent(tr.verticalContent) {
		PrintTextEnd(
			PrintTextParams{
				drawX:         bounding.OffsetBottomRight.X,
				drawY:         tr.computeCenterY(SplitTextLength) - 1,
				boundEndY:     bounding.OffsetBottomRight.Y,
				boundStartY:   bounding.OffsetTopLeft.Y,
				fg:            termbox.ColorWhite,
				bg:            termbox.ColorBlack,
				splitText:     tr.element.Text.SplitText,
				textDirection: 1,
			},
		)
	}

	if isBottomVerticalContent(tr.verticalContent) {
		PrintTextEnd(
			PrintTextParams{
				drawX:         bounding.OffsetBottomRight.X,
				drawY:         bounding.OffsetBottomRight.Y - SplitTextLength - 1,
				boundEndY:     bounding.OffsetBottomRight.Y,
				boundStartY:   bounding.OffsetTopLeft.Y,
				fg:            termbox.ColorWhite,
				bg:            termbox.ColorBlack,
				splitText:     tr.element.Text.SplitText,
				textDirection: 1,
			},
		)
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

			if y <= bounding.OffsetTopLeft.Y {
				break
			}

			centerByX := tr.computeCenterX(len(splitText) - 1)

			PrintRowStart(centerByX, y, termbox.ColorWhite, termbox.ColorBlack, splitText)
		}

	}

	if isCenterVerticalContent(tr.verticalContent) {
		centerPositionY := tr.computeCenterY(SplitTextLength - 1)

		for textIndex, splitText := range SplitText {
			y := centerPositionY + textIndex

			if y >= bounding.OffsetBottomRight.Y {
				break
			}

			if y <= bounding.OffsetTopLeft.Y {
				break
			}

			startPositionX := tr.computeCenterX(len(splitText) - 1)

			PrintRowStart(startPositionX, y, termbox.ColorWhite, termbox.ColorBlack, splitText)
		}
	}

	if isBottomVerticalContent(tr.verticalContent) {

		for textIndex, splitText := range SplitText {
			y := bounding.OffsetBottomRight.Y - SplitTextLength + textIndex

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			if y <= bounding.OffsetTopLeft.Y {
				continue
			}

			centerByX := tr.computeCenterX(len(splitText) - 1)

			PrintRowStart(centerByX, y, termbox.ColorWhite, termbox.ColorBlack, splitText)
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
	return (tr.element.Bounding.OffsetTopLeft.X + ((tr.element.Bounding.Width - tr.element.GetWidthOffset()) / 2)) - textLength/2
}

func (tr *textRenderer) computeCenterY(rows int) int {
	return (tr.element.Bounding.OffsetTopLeft.Y+tr.element.Bounding.OffsetBottomRight.Y)/2 - rows/2
}
