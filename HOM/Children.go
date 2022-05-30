package HOM

type Children struct {
	Elements []*Element
}

func (children *Children) GetMaxWidth(contentDirection ContentDirection) int {
	max := 0

	if children == nil {
		return max
	}

	if contentDirection == HorizontalDirection {

		for index, element := range children.Elements {
			max += element.Style.Width

			if index != 0 {
				max += element.getBorderOffset()
			}
		}
	} else {
		for _, element := range children.Elements {
			if max < element.Style.Width {
				max = element.Style.Width
			}
		}
	}

	return max
}

func (children *Children) GetMaxHeight(contentDirection ContentDirection) int {
	max := 0

	if children == nil {
		return max
	}

	if contentDirection == HorizontalDirection {
		for _, element := range children.Elements {
			if max < element.Style.Height {
				max = element.Style.Height

			}
		}
	} else {
		for index, element := range children.Elements {
			max += element.Style.Height

			if index != 0 {
				max += element.getBorderOffset()
			}
		}
	}

	return max
}
