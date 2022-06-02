package HOMF

import (
	"awesomeProject/consoleRenderer"
	"awesomeProject/pkg/HOM"
	"github.com/nsf/termbox-go"
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

	termbox.Flush()

}

func (homf *HOMFramework) Run() {
mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		case termbox.EventMouse:
			homf.propagateClick(ev.MouseX, ev.MouseY)
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func (homf *HOMFramework) FindElementById(element *HOM.Element, id string) *HOM.Element {

	if element.Children != nil {

		for _, elem := range element.Children.Elements {

			if elem.Style.ID == id {
				return elem
			}

			el := homf.FindElementById(elem, id)

			if el != nil {
				return el
			}
		}
	}

	return nil
}

func (homf *HOMFramework) propagateClick(MouseX, MouseY int) {
	homf.propagateClickRecursive(homf.Element, MouseX, MouseY)
}

func (homf *HOMFramework) propagateClickRecursive(Element *HOM.Element, MouseX, MouseY int) {

	if Element.Children != nil {
		for _, elem := range Element.Children.Elements {
			if homf.isInterceptElement(elem, MouseX, MouseY) {
				if elem.OnClick != nil {
					elem.OnClick(elem)
				}
			}

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
	homf.handOfMidas.PreprocessTree(homf.Element)

	consoleRenderer.RenderElement(homf.Element)

	termbox.Flush()
}
