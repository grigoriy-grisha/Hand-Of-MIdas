package Dom

type AlignContent int8
type VerticalContent int8

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

type Element struct {
	Style    Style
	Text     string
	Children *Element
}

func NewDomElement(Style Style, Text string, Children *Element) Element {
	return Element{Style, Text, Children}
}

func (e *Element) GetSize() (int, int) {
	return e.Style.Width, e.Style.Height
}

func (e *Element) getPosition() {

}
