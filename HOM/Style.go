package HOM

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
	Border          bool
	PaddingTop      int
	PaddingBottom   int
	PaddingLeft     int
	PaddingRight    int
	Width           int
	Height          int
	Y               int
	X               int
	VerticalContent VerticalContent
	AlignContent    AlignContent
}
