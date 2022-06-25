package HOMF

import (
	"awesomeProject/pkg/HOM"
	"awesomeProject/pkg/consoleRenderer"
	"github.com/nsf/termbox-go"
	"log"
)

type HOMFramework struct {
	Element     *HOM.Element
	handOfMidas *HOM.HandOfMidas
}

func NewHOMFramework(width int, height int) *HOMFramework {
	return &HOMFramework{handOfMidas: HOM.NewHandOfMidas(width, height)}
}

func (homf *HOMFramework) Init() func() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	return func() {
		termbox.Close()
	}
}

func (homf *HOMFramework) Mount(Element *HOM.Element) {
	homf.Element = Element
	homf.handOfMidas.PreprocessTree(Element)

	consoleRenderer.RenderElement(Element)
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	err := termbox.Flush()

	if err != nil {
		log.Fatal(err)
	}

}

var MouseDown termbox.Key = 65509
var LeftMouseUp termbox.Key = 65512

func (homf *HOMFramework) GetDetectClick() func(key termbox.Key, MouseX, MouseY int) {
	var clickKey termbox.Key = 0

	return func(key termbox.Key, MouseX, MouseY int) {
		if key == MouseDown {
			if clickKey == LeftMouseUp {
				homf.propagateClick(MouseX, MouseY)
			}
		}

		clickKey = key
	}
}

func (homf *HOMFramework) Run() {
	detectClick := homf.GetDetectClick()
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		case termbox.EventMouse:

			detectClick(ev.Key, ev.MouseX, ev.MouseY)
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func (homf *HOMFramework) FindElementById(element *HOM.Element, id string) *HOM.Element {
	if element.ID == id {
		return element
	}

	if element.Children == nil {
		return nil
	}

	for _, elem := range element.Children.Elements {
		el := homf.FindElementById(elem, id)

		if el != nil {
			return el
		}
	}

	return nil
}

func (homf *HOMFramework) propagateClick(MouseX, MouseY int) {
	homf.propagateClickRecursive(homf.Element, MouseX, MouseY)
}

func (homf *HOMFramework) propagateClickRecursive(Element *HOM.Element, MouseX, MouseY int) {
	if homf.isInterceptElement(Element, MouseX, MouseY) {
		if Element.OnClick != nil {
			Element.OnClick(Element)
		}
	}

	if Element.Children != nil {
		for _, elem := range Element.Children.Elements {

			homf.propagateClickRecursive(elem, MouseX, MouseY)
		}
	}

}

func (homf *HOMFramework) isInterceptElement(Element *HOM.Element, MouseX, MouseY int) bool {
	if MouseX >= Element.Bounding.ClientTopLeft.X && MouseY >= Element.Bounding.ClientTopLeft.Y {
		if MouseX >= Element.Bounding.ClientBottomLeft.X && MouseY <= Element.Bounding.ClientBottomLeft.Y {
			if MouseX <= Element.Bounding.ClientTopRight.X && MouseY >= Element.Bounding.ClientTopRight.Y {
				if MouseX <= Element.Bounding.ClientBottomRight.X && MouseY <= Element.Bounding.ClientBottomRight.Y {
					return true
				}
			}
		}
	}

	return false
}

func (homf *HOMFramework) Flush() {
	err := termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	if err != nil {
		return
	}

	homf.handOfMidas.PreprocessTree(homf.Element)
	consoleRenderer.RenderElement(homf.Element)

	termbox.Flush()
}
