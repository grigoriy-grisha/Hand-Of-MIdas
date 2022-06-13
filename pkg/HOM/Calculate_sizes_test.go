package HOM

import "testing"

var size = 100

func TestCalculateWidths(t *testing.T) {
	longText := "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! 	Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero! 	Lorem ipsum dolor sit amet, consectetur adipisicing elit. Amet ducimus inventore\nipsam obcaecati porro quas quia quos, saepe sapiente vero!"
	splitedLongText := splitLongText(size, longText)
	text := "hello"
	textLength := len(text)
	splitedLongTextLenght := len(splitedLongText)

	t.Run("calculate nil values width", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, testElement.getWidthOffset())
	})

	t.Run("calculate text width", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Text: &Text{Value: text}})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, textLength+testElement.getSizeOffset())
	})

	t.Run("calculate overflow text width", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Text: &Text{Value: longText}})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, size)
	})

	t.Run("calculate paddings", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Style: &Style{PaddingLeft: 5, PaddingRight: 5}})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, 10+testElement.getSizeOffset())
	})

	t.Run("calculate paddings width text", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Style: &Style{PaddingLeft: 5, PaddingRight: 5}, Text: &Text{Value: text}})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, 10+textLength+testElement.getSizeOffset())
	})

	t.Run("calculate width with children element with text", func(t *testing.T) {
		testElement := NewHOMElement(
			NewElementParams{
				Style:    &Style{PaddingLeft: 5, PaddingRight: 5},
				Children: &Children{Elements: []*Element{NewHOMElement(NewElementParams{Text: &Text{Value: text}})}},
			},
		)
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, 10+textLength+testElement.getSizeOffset())
	})

	t.Run("calculate width with children element text and padding", func(t *testing.T) {
		testElement := NewHOMElement(
			NewElementParams{
				Style: &Style{PaddingLeft: 5, PaddingRight: 5},
				Children: &Children{
					Elements: []*Element{
						NewHOMElement(
							NewElementParams{
								Style: &Style{PaddingLeft: 5, PaddingRight: 5},
								Text:  &Text{Value: text},
							},
						),
					},
				},
			},
		)
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, 20+textLength+testElement.getSizeOffset())
	})

	t.Run("calculate width with more children", func(t *testing.T) {
		testElement := NewHOMElement(
			NewElementParams{
				Children: &Children{Elements: []*Element{
					NewHOMElement(
						NewElementParams{
							Style: &Style{PaddingLeft: 5, PaddingRight: 5},
							Children: &Children{
								Elements: []*Element{
									NewHOMElement(NewElementParams{Style: &Style{PaddingLeft: 5, PaddingRight: 5}, Text: &Text{Value: text}}),
									NewHOMElement(NewElementParams{Style: &Style{PaddingLeft: 5, PaddingRight: 5}, Text: &Text{Value: text}}),
									NewHOMElement(NewElementParams{Style: &Style{PaddingLeft: 5, PaddingRight: 5}, Text: &Text{Value: text}}),
									NewHOMElement(NewElementParams{Text: &Text{Value: text}}),
								},
							},
						},
					)},
				},
			},
		)
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		// 5 + 5 - внешний блок = 10
		// 5 + 5 + 5 + 5 + 5 + 5 - паддинги для текстовых элементов = 30
		// 5 + 5 + 5 + 5 - длины всех строк = 20
		// 4 - коэфицент для текста  getSizeOffset = 4
		// итого 64
		assert(t, testElement.Bounding.Width, 64)
	})

	t.Run("calculate nil values height", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Height, testElement.getSizeOffset())
	})

	t.Run("calculate text height", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Text: &Text{Value: text}})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Height, 1+testElement.getSizeOffset())
	})

	t.Run("calculate long text height", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Text: &Text{Value: text}})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Height, splitedLongTextLenght+testElement.getSizeOffset())
	})

	t.Run("calculate nil values height width paddings", func(t *testing.T) {
		testElement := NewHOMElement(
			NewElementParams{
				Style: &Style{PaddingTop: 5, PaddingBottom: 5},
				Text:  &Text{Value: text},
			},
		)
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Height, 10+testElement.getSizeOffset())
	})

	t.Run("calculate nil values height width paddings", func(t *testing.T) {
		testElement := NewHOMElement(
			NewElementParams{
				Style: &Style{
					PaddingTop:    5,
					PaddingBottom: 5,
				},
				Text: &Text{Value: text},
			})
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Height, 10+1+testElement.getSizeOffset())
	})

	t.Run("calculate height with children element text and padding", func(t *testing.T) {
		testElement := NewHOMElement(
			NewElementParams{
				Style: &Style{PaddingTop: 5, PaddingBottom: 5},
				Children: &Children{
					Elements: []*Element{
						NewHOMElement(
							NewElementParams{
								Style: &Style{PaddingTop: 5, PaddingBottom: 5},
								Text:  &Text{Value: text},
							}),
					},
				},
			},
		)
		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Height, 20+1+testElement.getSizeOffset())
	})

	t.Run("calculate height with more children", func(t *testing.T) {
		testElement := NewHOMElement(
			NewElementParams{
				Children: &Children{Elements: []*Element{
					NewHOMElement(
						NewElementParams{
							Style: &Style{PaddingTop: 5, PaddingBottom: 5},
							Children: &Children{
								Elements: []*Element{
									NewHOMElement(NewElementParams{Style: &Style{PaddingTop: 5, PaddingBottom: 5}, Text: &Text{Value: text}}),
									NewHOMElement(NewElementParams{Style: &Style{PaddingTop: 5, PaddingBottom: 5}, Text: &Text{Value: text}}),
									NewHOMElement(NewElementParams{Style: &Style{PaddingTop: 5, PaddingBottom: 5}, Text: &Text{Value: text}}),
									NewHOMElement(NewElementParams{Text: &Text{Value: text}}),
								},
							},
						},
					),
				},
				},
			},
		)

		testHOM := NewHandOfMidas(size, size)

		testHOM.calculateSizes(size, size, testElement)

		assert(t, testElement.Bounding.Width, 24)
	})
}
