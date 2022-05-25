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

func RenderElement(element *HOM.Element) {
	//tood вчислять в в препроцесс
	topX := element.Style.X
	topY := element.Style.Y

	topRightX := topX + element.Style.Width
	bottomLeftY := topY + element.Style.Height
	bottomRightX := element.Style.X + element.Style.Width

	termbox.SetCell(topX, topY, borderTopLeft, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(topRightX, topY, borderTopRight, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(topX, bottomLeftY, borderBotomLeft, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(bottomRightX, bottomLeftY, borderBottomRight, termbox.ColorWhite, termbox.ColorBlack)

	for i := topY + 1; i < bottomLeftY; i++ {
		termbox.SetCell(topX, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(bottomRightX, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
	}

	for i := topX + 1; i < bottomRightX; i++ {
		termbox.SetCell(i, topY, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, bottomLeftY, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
	}

	if TextIsNotEmpty(element.Text.Value) {
		textRenderer := NewTextRenderer(
			NewTextRendererParams{
				alignContent:    element.Style.AlignContent,
				verticalContent: element.Style.VerticalContent,
				width:           element.Style.Width,
				height:          element.Style.Height,
				topX:            topX,
				topY:            topY,
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
