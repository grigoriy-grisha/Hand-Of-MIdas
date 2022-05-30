package HOM

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

	Element.computeBounding()

	//todo, возможно, это не должно тут быть
	normalizedWidth := hom.Window.Width - Element.Style.PaddingLeft - Element.Style.PaddingRight + Element.getSizeOffset()
	normalizedHeight := hom.Window.Height - Element.Style.PaddingTop - Element.Style.PaddingBottom + Element.getSizeOffset()

	if Element.Text != nil {
		Element.Text.CalculateTextHyphens(
			hom.Window.Width,
			Element.Style.PaddingLeft+Element.Style.PaddingRight+Element.getSizeOffset(),
		)
	}

	maxWidth := normalizedWidth / len(Element.Children.Elements)

	hom.calculateLayout(maxWidth, normalizedHeight, Element.Bounding.OffsetTopLeft, Element)
}

// todo + 1 коэфицент это из-за border
func (hom *HandOfMidas) calculateLayout(parentWidth int, parentHeight int, coords *Coords, Element *Element) {
	prevCoords := &Coords{X: coords.X, Y: coords.Y}

	for _, element := range Element.Children.Elements {
		element.Style.X = prevCoords.X
		element.Style.Y = prevCoords.Y

		element.ComputeElementSize(parentWidth, parentHeight)

		if element.Children != nil {
			hom.calculateLayout(
				element.getAvailableWidth(parentWidth),
				element.getAvailableHeight(Element.Style.ContentDirection, parentHeight),
				element.getCoordsForNextElement(Element.Style.ContentDirection, prevCoords),
				element,
			)

			element.Style.Width = element.getWidthWithChildren(element.Style.ContentDirection, parentWidth)
			element.Style.Height = element.getHeightWithChildren(element.Style.ContentDirection, parentHeight)
		}

		element.computeBounding()

		if Element.Style.ContentDirection == HorizontalDirection {
			prevCoords = &Coords{
				X: element.Bounding.ClientTopRight.X + element.getBorderOffset(),
				Y: element.Bounding.ClientTopRight.Y,
			}
		} else {
			prevCoords = &Coords{
				X: element.Bounding.ClientBottomLeft.X,
				Y: element.Bounding.ClientBottomLeft.Y + element.getBorderOffset(),
			}
		}

	}
}
