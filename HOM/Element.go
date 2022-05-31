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
	Style    *Style
	Text     *Text
	Children *Children
	Bounding *Bounding
}

func NewHOMElement(Style *Style, Text *Text, Children *Children) *Element {
	element := &Element{Style: Style, Text: Text, Children: Children}
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
	return element.Text.ValueLength +
		element.Style.PaddingLeft +
		element.Style.PaddingRight +
		element.getSizeOffset()
}

func (element *Element) computeWidth(parentWidth int) int {
	if element.Text.SplitTextLength > 1 {
		return parentWidth
	}

	computedWidth := element.getElementWidth()
	if computedWidth > parentWidth {
		return parentWidth
	}

	return computedWidth
}

func (element *Element) getElementHeight() int {
	return element.Text.SplitTextLength +
		element.Style.PaddingTop +
		element.Style.PaddingBottom +
		element.getSizeOffset()
}

func (element *Element) computeHeight(parentHeight int) int {
	computedHeight := element.getElementHeight()

	if computedHeight > parentHeight {
		return parentHeight
	}

	return computedHeight
}

func (element *Element) ComputeElementSize(parentWidth, parentHeight int) {
	if element.Text != nil {
		element.Text.CalculateTextHyphens(parentWidth, element.getWidthOffset())
		element.Bounding.Width = element.computeWidth(parentWidth)
		element.Bounding.Height = element.computeHeight(parentHeight)
	}
}

func (element *Element) getAvailableWidth(parentWidth int) int {
	return (parentWidth -
		element.Style.PaddingLeft -
		element.Style.PaddingRight) /
		len(element.Children.Elements)
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

func (element *Element) setWidth(parentWidth int) {
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
