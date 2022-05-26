package HOM

import (
	"strings"
)

type AlignContent int8
type VerticalContent int8

//todo вынести в отдельный модуль стилей
var (
	AlignContentLeft   AlignContent = 0
	AlignContentCenter AlignContent = 1
	AlignContentRight  AlignContent = 2
)

var (
	VerticalContentTop    VerticalContent = 0
	VerticalContentCenter VerticalContent = 1
	VerticalContentBottom VerticalContent = 2
)

type Style struct {
	PaddingTop      int
	PaddingBottom   int
	PaddingLeft     int
	PaddingRight    int
	Width           int
	Height          int
	Y               int
	X               int
	AlignContent    AlignContent
	VerticalContent VerticalContent
}

type CompositeElement interface {
	GetSize() (int, int)
}

type Children struct {
	Elements []*Element
}

type Text struct {
	Value     string
	SplitText []string
}

type Element struct {
	Style    *Style
	Text     *Text
	Children *Children
	Parent   *Element
}

func NewDomElement(Style *Style, Text *Text, Children *Children) *Element {
	return &Element{Style: Style, Text: Text, Children: Children}
}

func (e *Element) GetSize() (int, int) {
	return e.Style.Width, e.Style.Height
}

func (e *Element) getPosition() {

}

type Window struct {
	Width   int
	Height  int
	Element *Element
}

type HandOfMidas struct {
	Window Window
}

func NewHandOfMidas(width int, height int) *HandOfMidas {
	handOfMidas := &HandOfMidas{}
	handOfMidas.SetSizeWindow(width, height)

	return handOfMidas
}

func (hom *HandOfMidas) SetSizeWindow(width int, height int) {
	hom.Window.Width = width
	hom.Window.Height = height
}

func (hom *HandOfMidas) PreprocessTree(Element *Element) {
	hom.Window.Element = Element

	if Element.Style.Width == 0 {
		Element.Style.Width = hom.Window.Width
	}

	if Element.Style.Height == 0 {
		Element.Style.Height = hom.Window.Height
	}

	//todo, возможно, это не должно тут быть
	normalizedWidth := hom.Window.Width - Element.Style.PaddingLeft - Element.Style.PaddingRight

	if Element.Text != nil {
		Element.Text.SplitText = SplitLongText(normalizedWidth, Element.Text.Value)
	}

	//Element.Style.Width = hom.Window.Width + Element.Style.PaddingLeft + Element.Style.PaddingRight
	//todo разобрать , где потеряна 1
	//Element.Style.Height = len(Element.Text.SplitText) + Element.Style.PaddingTop + Element.Style.PaddingBottom + 1

	//hom.calculateLayout(hom.Window.Width, hom.Window.Height, Element)
}

// todo разделить по строкам текст, рефткоринг

func SplitLongText(width int, text string) []string {
	var splitText []string

	splitStrings := strings.Split(text, " ")

	preparedString := ""

	for index, splitString := range splitStrings {
		preparedStringLength := len(preparedString + splitString)

		if width <= preparedStringLength {
			splitText = append(splitText, preparedString)
			preparedString = ""
			preparedStringLength = 0
		}

		if preparedStringLength == 0 || index == 0 {
			preparedString += splitString
		} else {
			preparedString += " " + splitString
		}
	}

	if len(preparedString) != 0 {
		splitText = append(splitText, preparedString)
	}

	return splitText
}

func (hom *HandOfMidas) calculateLayout(parentWidth int, parentHeight int, Element *Element) {
	if Element.Children == nil {
		return
	}

	//todo Тут надо применять flex свойства
	for _, elem := range Element.Children.Elements {
		//elem.Parent = Element

		if elem.Style.Height == 0 {
			elem.Style.Height = parentHeight
		}

		if elem.Style.Width == 0 {
			elem.Style.Width = parentWidth
		}

		hom.calculateLayout(Element.Style.Width, Element.Style.Height, elem)
	}
}
