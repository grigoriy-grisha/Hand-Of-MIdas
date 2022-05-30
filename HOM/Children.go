package HOM

type Children struct {
	Elements []*Element
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
