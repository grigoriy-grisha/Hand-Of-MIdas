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
	Border          bool
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

type Coords struct {
	X int
	Y int
}

type Bounding struct {
	ClientTopLeft     *Coords
	ClientBottomLeft  *Coords
	ClientTopRight    *Coords
	ClientBottomRight *Coords
	OffsetTopLeft     *Coords
	OffsetBottomLeft  *Coords
	OffsetTopRight    *Coords
	OffsetBottomRight *Coords
}

type Element struct {
	Style    *Style
	Text     *Text
	Children *Children
	Parent   *Element
	Bounding *Bounding
}

func NewDomElement(Style *Style, Text *Text, Children *Children) *Element {
	element := &Element{Style: Style, Text: Text, Children: Children}
	element.Bounding = &Bounding{}

	return element
}

func (e *Element) GetSize() (int, int) {
	return e.Style.Width, e.Style.Height
}

type Window struct {
	Width   int
	Height  int
	Element *Element
}

type HandOfMidas struct {
	Window Window
}

func (children *Children) GetMaxWidth() int {
	max := 0

	if children == nil {
		return max
	}

	for _, element := range children.Elements {
		//todo 1 это отступ между элементами
		max += element.Style.Width + 1
	}

	return max - 1
}

func (children *Children) GetMaxHeight() int {
	max := 0

	if children == nil {
		return max
	}
	for _, element := range children.Elements {
		if max < element.Style.Height {
			max = element.Style.Height
		}
	}

	return max
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

func (hom *HandOfMidas) getBorderOffset(Element *Element) int {
	if Element.Style.Border {
		return 1
	}
	return 1
}

func (hom *HandOfMidas) computeBounding(Element *Element) {
	ClientX := Element.Style.X
	ClientY := Element.Style.Y
	FullClientY := ClientY + Element.Style.Height
	FullClientX := ClientX + Element.Style.Width

	Element.Bounding.ClientTopLeft = &Coords{X: ClientX, Y: ClientY}
	Element.Bounding.ClientBottomLeft = &Coords{X: ClientX, Y: FullClientY}
	Element.Bounding.ClientTopRight = &Coords{X: FullClientX, Y: ClientY}
	Element.Bounding.ClientBottomRight = &Coords{X: FullClientX, Y: FullClientY}

	OffsetY := ClientY + Element.Style.PaddingTop
	OffsetX := ClientX + Element.Style.PaddingLeft
	FullOffsetY := FullClientY - Element.Style.PaddingBottom
	FullOffsetX := FullClientX - Element.Style.PaddingRight

	Element.Bounding.OffsetTopLeft = &Coords{X: OffsetX, Y: OffsetY}
	Element.Bounding.OffsetBottomLeft = &Coords{X: OffsetX, Y: FullOffsetY}
	Element.Bounding.OffsetTopRight = &Coords{X: FullOffsetX, Y: OffsetY}
	Element.Bounding.OffsetBottomRight = &Coords{X: FullOffsetX, Y: FullOffsetY}

}

func (hom *HandOfMidas) PreprocessTree(Element *Element) {
	hom.Window.Element = Element

	if Element.Style.Width == 0 {
		Element.Style.Width = hom.Window.Width
	}

	if Element.Style.Height == 0 {
		Element.Style.Height = hom.Window.Height
	}

	hom.computeBounding(Element)

	//todo, возможно, это не должно тут быть
	normalizedWidth := hom.Window.Width - Element.Style.PaddingLeft - Element.Style.PaddingRight + hom.getBorderOffset(Element)
	normalizedHeight := hom.Window.Height - Element.Style.PaddingTop - Element.Style.PaddingBottom + hom.getBorderOffset(Element)

	if Element.Text != nil {
		Element.Text.SplitText = SplitLongText(normalizedWidth, Element.Text.Value)
	}

	maxWidth := normalizedWidth / len(Element.Children.Elements)

	hom.calculateLayout(maxWidth, normalizedHeight, Element.Bounding.OffsetTopLeft, Element)
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

// todo + 1 коэфицент это из-за border
func (hom *HandOfMidas) calculateLayout(parentWidth int, parentHeight int, coords *Coords, Element *Element) {
	prevCoords := &Coords{
		X: coords.X,
		Y: coords.Y,
	}

	for _, element := range Element.Children.Elements {
		bodrderOffset := hom.getBorderOffset(element)
		element.Style.X = prevCoords.X
		element.Style.Y = prevCoords.Y

		if element.Text != nil {
			element.Text.SplitText = SplitLongText(
				parentWidth-element.Style.PaddingLeft-element.Style.PaddingRight,
				element.Text.Value,
			)
		}

		if element.Text != nil {
			if len(element.Text.SplitText) > 1 {
				element.Style.Width = parentWidth
			} else {
				computedWidth := len(element.Text.Value) + element.Style.PaddingLeft + element.Style.PaddingRight + bodrderOffset

				if computedWidth > parentWidth {
					element.Style.Width = parentWidth
				} else {
					element.Style.Width = computedWidth
				}
			}

			computedHeight := len(element.Text.SplitText) +
				element.Style.PaddingTop +
				element.Style.PaddingBottom + bodrderOffset

			if computedHeight > parentHeight {
				element.Style.Height = parentHeight
			} else {
				element.Style.Height = computedHeight
			}
		}

		if element.Children != nil {
			hom.calculateLayout(
				(parentWidth-element.Style.PaddingLeft-element.Style.PaddingRight)/len(element.Children.Elements),
				parentHeight-element.Style.PaddingTop-element.Style.PaddingBottom,
				&Coords{X: prevCoords.X + element.Style.PaddingLeft, Y: prevCoords.Y + element.Style.PaddingTop},
				element,
			)

			computedWidth := element.Children.GetMaxWidth() + element.Style.PaddingLeft + element.Style.PaddingRight
			computedHeight := element.Children.GetMaxHeight() + element.Style.PaddingTop + element.Style.PaddingBottom

			if computedWidth > parentWidth {
				element.Style.Width = parentWidth
			} else {
				element.Style.Width = computedWidth
			}

			if computedHeight > parentHeight {
				element.Style.Height = parentHeight
			} else {
				element.Style.Height = computedHeight
			}
		}

		hom.computeBounding(element)
		prevCoords = &Coords{
			X: element.Bounding.ClientTopRight.X + 1,
			Y: element.Bounding.ClientTopRight.Y,
		}

	}
}
