package HOM

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
	Width             int
	Height            int
}

type Element struct {
	ID             string
	Style          *Style
	Text           *Text
	Children       *Children
	Bounding       *Bounding
	ParentBounding *Bounding
	OnClick        func(element *Element)
}

type NewElementParams struct {
	ID       string
	Style    *Style
	Text     *Text
	Children *Children
	OnClick  func(element *Element)
}

func NewHOMElement(params NewElementParams) *Element {
	style := params.Style

	if params.Style == nil {
		style = &Style{}
	}

	element := &Element{ID: params.ID, Style: style, Text: params.Text, Children: params.Children, OnClick: params.OnClick}
	element.Bounding = &Bounding{}

	return element
}

func (element *Element) getBorderOffset() int {
	if element.Style.Border {
		return 1
	}
	return 0
}

func (element *Element) getSizeOffset() int {
	return 1
}

func (element *Element) getWidthOffset() int {
	return element.Style.PaddingLeft +
		element.Style.PaddingRight +
		element.getSizeOffset()
}

func (element *Element) getElementWidth() int {
	valueLength := 0

	if element.Text != nil {
		valueLength = element.Text.ValueLength
	}

	return valueLength +
		element.Style.PaddingLeft +
		element.Style.PaddingRight +
		element.getSizeOffset()
}

func (element *Element) computeWidth(parentWidth int) int {
	computedWidth := element.getElementWidth()

	if element.Text == nil {
		return computedWidth
	}

	if element.Text.SplitTextLength > 1 {
		return parentWidth
	}

	if computedWidth > parentWidth {
		return parentWidth
	}

	return computedWidth
}

func (element *Element) getElementHeight() int {
	splitTextLength := 0

	if element.Text != nil {
		splitTextLength = element.Text.SplitTextLength
	}

	return splitTextLength +
		element.Style.PaddingTop +
		element.Style.PaddingBottom +
		element.getSizeOffset()
}

func (element *Element) computeHeight(parentHeight int) int {
	computedHeight := element.getElementHeight()

	if element.Text == nil {
		return computedHeight
	}

	if computedHeight > parentHeight {
		return parentHeight
	}

	return computedHeight
}

func (element *Element) ComputeElementSize(parentWidth, parentHeight int) {
	if element.Text != nil {
		element.Text.CalculateTextHyphens(parentWidth, element.getWidthOffset())
	}

	if element.Bounding.Width == 0 {
		element.Bounding.Width = element.computeWidth(parentWidth)
	}

	if element.Bounding.Height == 0 {
		element.Bounding.Height = element.computeHeight(parentHeight)
	}
}

func (element *Element) getAvailableWidth(parentWidth int) int {
	return parentWidth -
		element.Style.PaddingLeft -
		element.Style.PaddingRight
}

func (element *Element) getAvailableHeight(contentDirection ContentDirection, parentHeight int) int {
	if contentDirection == HorizontalDirection {
		return parentHeight -
			element.Style.PaddingTop -
			element.Style.PaddingBottom
	}

	computedHeight := element.Children.GetMaxHeight(contentDirection) +
		element.Style.PaddingLeft +
		element.Style.PaddingRight

	if computedHeight > parentHeight {
		return parentHeight
	}

	return computedHeight

}

func (element *Element) getCoordsForNextElement(ContentDirection ContentDirection, prevCoords *Coords) *Coords {
	if ContentDirection == HorizontalDirection {
		return &Coords{
			X: prevCoords.X + element.Style.PaddingLeft,
			Y: prevCoords.Y + element.Style.PaddingTop,
		}
	}

	return &Coords{
		X: prevCoords.X + element.Style.PaddingLeft,
		Y: prevCoords.Y + element.Style.PaddingBottom,
	}
}

func (element *Element) getWidthWithChildren(contentDirection ContentDirection, parentWidth int) int {
	computedWidth := element.Children.GetMaxWidth(contentDirection) +
		element.Style.PaddingLeft +
		element.Style.PaddingRight

	if computedWidth > parentWidth {
		return parentWidth
	}

	return computedWidth
}

func (element *Element) getHeightWithChildren(contentDirection ContentDirection, parentHeight int) int {
	computedHeight := element.Children.GetMaxHeight(contentDirection) +
		element.Style.PaddingTop +
		element.Style.PaddingBottom

	if computedHeight > parentHeight {
		return parentHeight
	}

	return computedHeight

}

//todo рефакторинг
func (element *Element) setWidth(parentWidth int) {

	if element.Style.Width == nil {
		element.Bounding.Width = element.getWidthWithChildren(element.Style.ContentDirection, parentWidth)
		return
	}

	convertedIntWidth, isInt := element.Style.Width.(int)

	if isInt {
		if convertedIntWidth == 0 {
			element.Bounding.Width = element.getWidthWithChildren(element.Style.ContentDirection, parentWidth)
			return
		}

		element.Bounding.Width = convertedIntWidth
		return
	}

	convertedStringWidth, err := parsePercentStringToPercentFloat(element.Style.Width)

	if err == nil {
		if convertedStringWidth == 0 {
			element.Bounding.Width = element.getWidthWithChildren(element.Style.ContentDirection, parentWidth)
			return
		}

		computedWidth := convertedStringWidth * float64(parentWidth)
		element.Bounding.Width = int(computedWidth)
	}
}

func (element *Element) setHeight(parentHeight int) {

	if element.Style.Height == nil {
		element.Bounding.Height = element.getHeightWithChildren(element.Style.ContentDirection, parentHeight)
		return
	}

	convertedIntHeight, isInt := element.Style.Height.(int)

	if isInt {
		if convertedIntHeight == 0 {
			element.Bounding.Height = element.getHeightWithChildren(element.Style.ContentDirection, parentHeight)
			return
		}

		element.Bounding.Height = convertedIntHeight
		return
	}

	convertedStringHeight, err := parsePercentStringToPercentFloat(element.Style.Height)

	if err == nil {
		if convertedStringHeight == 0 {
			element.Bounding.Height = element.getHeightWithChildren(element.Style.ContentDirection, parentHeight)
			return
		}

		computedWidth := convertedStringHeight * float64(parentHeight)
		element.Bounding.Height = int(computedWidth)
	}
}

func (element *Element) computeBounding() {
	ClientX := element.Style.X
	ClientY := element.Style.Y
	FullClientY := ClientY + element.Bounding.Height
	FullClientX := ClientX + element.Bounding.Width

	element.Bounding.ClientTopLeft = &Coords{X: ClientX, Y: ClientY}
	element.Bounding.ClientBottomLeft = &Coords{X: ClientX, Y: FullClientY}
	element.Bounding.ClientTopRight = &Coords{X: FullClientX, Y: ClientY}
	element.Bounding.ClientBottomRight = &Coords{X: FullClientX, Y: FullClientY}

	OffsetY := ClientY + element.Style.PaddingTop
	OffsetX := ClientX + element.Style.PaddingLeft
	FullOffsetY := FullClientY - element.Style.PaddingBottom
	FullOffsetX := FullClientX - element.Style.PaddingRight

	element.Bounding.OffsetTopLeft = &Coords{X: OffsetX, Y: OffsetY}
	element.Bounding.OffsetBottomLeft = &Coords{X: OffsetX, Y: FullOffsetY}
	element.Bounding.OffsetTopRight = &Coords{X: FullOffsetX, Y: OffsetY}
	element.Bounding.OffsetBottomRight = &Coords{X: FullOffsetX, Y: FullOffsetY}

}
