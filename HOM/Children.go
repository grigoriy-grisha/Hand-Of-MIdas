package HOM

type Children struct {
	Elements []*Element
}

func (children *Children) GetMaxWidth() int {
	max := 0

	if children == nil {
		return max
	}

	for index, element := range children.Elements {
		max += element.Style.Width

		if index != 0 {
			max += element.getBorderOffset()
		}
	}

	return max
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
