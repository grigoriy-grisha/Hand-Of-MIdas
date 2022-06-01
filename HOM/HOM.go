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

func (hom *HandOfMidas) calculateSizes(parentWidth int, parentHeight int, Element *Element) {
	Element.ComputeElementSize(parentWidth, parentHeight)

	if Element.Children == nil {
		return
	}

	for _, element := range Element.Children.Elements {

		if element.Children != nil {
			availableWidth := element.getAvailableWidth(parentWidth)

			hom.calculateSizes(
				availableWidth,
				//TODO УЗНАТЬ почему Height работает не правильно
				parentHeight,
				element,
			)
		}

		if element.Text != nil {
			element.ComputeElementSize(parentWidth, parentHeight)
		}

	}

	Element.setWidth(parentWidth)
	Element.setHeight(parentHeight)
}

func (hom *HandOfMidas) PreprocessTree(Element *Element) {
	hom.Window.Element = Element

	if Element.Bounding.Width == 0 {
		Element.Bounding.Width = hom.Window.Width
	}

	if Element.Bounding.Height == 0 {
		Element.Bounding.Height = hom.Window.Height
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

	hom.calculateSizes(normalizedWidth, normalizedHeight, Element)
	hom.calculateLayout(Element.Bounding.OffsetTopLeft, Element)
}

// todo + 1 коэфицент это из-за border
func (hom *HandOfMidas) calculateLayout(coords *Coords, Element *Element) {
	var prevCoords *Coords = nil

	if Element.Style.AlignItems == AlignItemsStart {
		prevCoords = &Coords{X: coords.X, Y: coords.Y}
	}

	if Element.Style.AlignItems == AlignItemsEnd {
		prevCoords = &Coords{X: coords.X + Element.Bounding.Width, Y: coords.Y}
	}

	for _, element := range Element.Children.Elements {
		element.ParentBounding = Element.Bounding
		element.Style.X = prevCoords.X
		element.Style.Y = prevCoords.Y

		if element.Children != nil {
			hom.calculateLayout(
				element.getCoordsForNextElement(Element.Style.ContentDirection, Element.Style.AlignContent, prevCoords),
				element,
			)
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
