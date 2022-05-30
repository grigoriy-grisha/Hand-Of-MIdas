package consoleRenderer

import (
	"awesomeProject/HOM"
	"github.com/nsf/termbox-go"
)

var borderTopLeft rune = 0x250C
var borderTopRight rune = 0x2510
var borderBottomLeft rune = 0x2514
var borderBottomRight rune = 0x2518
var borderHorizontal rune = 0x2502
var borderVertical rune = 0x2500

func selectCell(coords *HOM.Coords, border rune) {
	termbox.SetCell(coords.X, coords.Y, border, termbox.ColorWhite, termbox.ColorBlack)
}

func drawBorder(element *HOM.Element) {
	bounding := element.Bounding

	selectCell(bounding.ClientTopLeft, borderTopLeft)
	selectCell(bounding.ClientBottomLeft, borderBottomLeft)
	selectCell(bounding.ClientTopRight, borderTopRight)
	selectCell(bounding.ClientBottomRight, borderBottomRight)

	for i := bounding.ClientTopLeft.Y + 1; i < bounding.ClientBottomLeft.Y; i++ {
		termbox.SetCell(bounding.ClientTopLeft.X, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(bounding.ClientBottomRight.X, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
	}

	for i := bounding.ClientTopLeft.X + 1; i < bounding.ClientBottomRight.X; i++ {
		termbox.SetCell(i, bounding.ClientTopLeft.Y, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, bounding.ClientBottomLeft.Y, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
	}
	//
	//selectCell(bounding.OffsetTopLeft, borderTopLeft)
	//selectCell(bounding.OffsetBottomLeft, borderBottomLeft)
	//selectCell(bounding.OffsetTopRight, borderTopRight)
	//selectCell(bounding.OffsetBottomRight, borderBottomRight)
	//
	//for i := bounding.OffsetTopLeft.Y + 1; i < bounding.OffsetBottomLeft.Y; i++ {
	//	termbox.SetCell(bounding.OffsetTopLeft.X, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
	//	termbox.SetCell(bounding.OffsetBottomRight.X, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
	//}
	//
	//for i := bounding.OffsetTopLeft.X + 1; i < bounding.OffsetBottomRight.X; i++ {
	//	termbox.SetCell(i, bounding.OffsetTopLeft.Y, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
	//	termbox.SetCell(i, bounding.OffsetBottomLeft.Y, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
	//}
}

func RenderElement(element *HOM.Element) {

	if element.Style.Border {
		drawBorder(element)
	}

	if TextIsNotEmpty(element.Text) {
		NewTextRenderer(
			NewTextRendererParams{
				alignContent:    element.Style.AlignContent,
				verticalContent: element.Style.VerticalContent,
				element:         element,
			},
		).renderText()
	}

	if element.Children == nil {
		return
	}

	for _, element := range element.Children.Elements {
		RenderElement(element)
	}

}
