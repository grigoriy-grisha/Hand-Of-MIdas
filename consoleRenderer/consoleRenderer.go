package consoleRenderer

import (
	"awesomeProject/Dom"
	"github.com/nsf/termbox-go"
)

var borderTopLeft rune = 0x250C
var borderTopRight rune = 0x2510
var borderBotomLeft rune = 0x2514
var borderBottomRight rune = 0x2518
var borderHorizontal rune = 0x2502
var borderVertical rune = 0x2500

// todo текст, который не вмещается вообще обрезать

func RenderElement(element Dom.Element) {
	topLeftX := element.Style.X
	topLeftY := element.Style.Y

	topRightX := topLeftX + element.Style.Width
	bottomLeftY := topLeftY + element.Style.Height
	bottomRightX := element.Style.X + element.Style.Width

	termbox.SetCell(topLeftX, topLeftY, borderTopLeft, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(topRightX, topLeftY, borderTopRight, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(topLeftX, bottomLeftY, borderBotomLeft, termbox.ColorWhite, termbox.ColorBlack)
	termbox.SetCell(bottomRightX, bottomLeftY, borderBottomRight, termbox.ColorWhite, termbox.ColorBlack)

	for i := topLeftY + 1; i < bottomLeftY; i++ {
		termbox.SetCell(topLeftX, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(bottomRightX, i, borderHorizontal, termbox.ColorWhite, termbox.ColorBlack)
	}

	for i := topLeftX + 1; i < bottomRightX; i++ {
		termbox.SetCell(i, topLeftY, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
		termbox.SetCell(i, bottomLeftY, borderVertical, termbox.ColorWhite, termbox.ColorBlack)
	}

	if TextIsNotEmpty(element.Text) {
		textRenderer := TextRenderer{alignContent: element.Style.AlignContent, verticalContent: element.Style.VerticalContent}
		textRenderer.renderText(element)
	}
}
