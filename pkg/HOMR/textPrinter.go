package HOMR

import "github.com/nsf/termbox-go"

type PrintTextParams struct {
	drawX         int
	drawY         int
	boundStartY   int
	boundEndY     int
	fg            termbox.Attribute
	bg            termbox.Attribute
	splitText     []string
	textDirection int
}

func PrintRowStart(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func PrintRowEnd(x, y int, fg, bg termbox.Attribute, text string) {
	textLength := len(text) - 1

	for i := textLength; i >= 0; i-- {
		x := x - 1 - i
		termbox.SetCell(x, y, rune(text[textLength-i]), fg, bg)
	}
}

func PrintTextStart(params PrintTextParams) {
	for textIndex, spitText := range params.splitText {
		y := params.drawY + textIndex + 1

		if y <= params.boundStartY {
			break
		}

		if y >= params.boundEndY {
			break
		}

		PrintRowStart(params.drawX, y, params.fg, params.bg, spitText)
	}
}

func PrintTextEnd(params PrintTextParams) {
	for textIndex, spitText := range params.splitText {
		y := params.drawY + textIndex + 1

		if y <= params.boundStartY {
			continue
		}

		if y >= params.boundEndY {
			continue
		}

		PrintRowEnd(params.drawX, y, params.fg, params.bg, spitText)
	}
}
