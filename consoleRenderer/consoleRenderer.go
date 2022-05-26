package consoleRenderer

import (
	"awesomeProject/HOM"
	"github.com/nsf/termbox-go"
)

var borderTopLeft rune = 0x250C
var borderTopRight rune = 0x2510
var borderBotomLeft rune = 0x2514
var borderBottomRight rune = 0x2518
var borderHorizontal rune = 0x2502
var borderVertical rune = 0x2500

// todo текст, который не вмещается вообще обрезать

func selectCell(coords *HOM.Coords, border rune) {
	termbox.SetCell(coords.X, coords.Y, border, termbox.ColorWhite, termbox.ColorBlack)
}

func RenderElement(element *HOM.Element) {
	bounding := element.Bounding

	selectCell(bounding.ClientTopLeft, borderTopLeft)
	selectCell(bounding.ClientBottomLeft, borderBotomLeft)
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

	if TextIsNotEmpty(element.Text.Value) {
		textRenderer := NewTextRenderer(
			NewTextRendererParams{
				alignContent:    element.Style.AlignContent,
				verticalContent: element.Style.VerticalContent,
				width:           element.Style.Width,
				height:          element.Style.Height,
				topX:            bounding.ClientTopLeft.X,
				topY:            bounding.ClientTopLeft.Y,
				paddingBottom:   element.Style.PaddingBottom,
				paddingTop:      element.Style.PaddingTop,
				paddingLeft:     element.Style.PaddingLeft,
				paddingRight:    element.Style.PaddingRight,
			})

		textRenderer.renderText(*element)
	}

	if element.Children != nil && len(element.Children.Elements) != 0 {
		for _, elem := range element.Children.Elements {
			RenderElement(elem)
		}
	}

}
