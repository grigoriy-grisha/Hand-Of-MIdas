package HOM

type AlignContent int8
type VerticalContent int8
type ContentDirection int8
type AlignItems int8

var (
	AlignItemsStart  AlignItems = 0
	AlignItemsCenter AlignItems = 1
	AlignItemsEnd    AlignItems = 2
)

var (
	HorizontalDirection ContentDirection = 0
	VerticalDirection   ContentDirection = 1
)

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
	Border           bool
	PaddingTop       int
	PaddingBottom    int
	PaddingLeft      int
	PaddingRight     int
	Y                int
	X                int
	ID               string
	ContentDirection ContentDirection
	VerticalContent  VerticalContent
	AlignContent     AlignContent
	AlignItems       AlignItems
	Height           interface{}
	Width            interface{}
}
