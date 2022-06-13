package consoleRenderer

import "github.com/nsf/termbox-go"

type PrintTextParams struct {
	drawX       int
	drawY       int
	boundStartY int
	boundEndY   int
	fg          termbox.Attribute
	bg          termbox.Attribute
	splitText   []string
}

func PrintRow(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func PrintText(params PrintTextParams) {
	for textIndex, spitText := range params.splitText {
		y := params.drawY + textIndex + 1

		if y <= params.boundStartY {
			break
		}

		if y >= params.boundEndY {
			break
		}

		PrintRow(params.drawX, y, params.fg, params.bg, spitText)
	}
}
