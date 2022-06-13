package HOM

import (
	"testing"
)

func TestElementSetWidth(t *testing.T) {
	t.Run("set int width 100", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Style: &Style{Width: 100}})
		testElement.setWidth(100)

		assert(t, testElement.Bounding.Width, 100)
	})

	t.Run("set int width 50", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Style: &Style{Width: 50}})
		testElement.setWidth(50)

		assert(t, testElement.Bounding.Width, 50)
	})

	t.Run("set percentage width", func(t *testing.T) {
		t.Run("width 100%", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Width: "100%"}})
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width, 100)
		})

		t.Run("width 75%", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Width: "75%"}})
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width, 75)
		})

		t.Run("width 25%", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Width: "25%"}})
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width, 25)
		})

		t.Run("broken string", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Width: "fuck you"}})
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width, 0)
		})

		t.Run("struct", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Width: Style{}}})
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width, 0)
		})

	})
}

func TestElementSetHeight(t *testing.T) {
	t.Run("set int Height 100", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Style: &Style{Height: 100}})
		testElement.setHeight(100)

		assert(t, testElement.Bounding.Height, 100)
	})

	t.Run("set int Height 50", func(t *testing.T) {
		testElement := NewHOMElement(NewElementParams{Style: &Style{Height: 50}})
		testElement.setHeight(50)

		assert(t, testElement.Bounding.Height, 50)
	})

	t.Run("set percentage Height", func(t *testing.T) {
		t.Run("Height 100%", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Height: "100%"}})
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height, 100)
		})

		t.Run("Height 75%", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Height: "75%"}})
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height, 75)
		})

		t.Run("Height 25%", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Height: "25%"}})
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height, 25)
		})

		t.Run("broken string", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Height: "fuck you"}})
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height, 0)
		})

		t.Run("struct", func(t *testing.T) {
			testElement := NewHOMElement(NewElementParams{Style: &Style{Height: Style{}}})
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height, 0)
		})
	})
}
