package HOM

import (
	"testing"
)

func TestElementSetWidth(t *testing.T) {
	t.Run("set int width 100", func(t *testing.T) {
		testElement := NewHOMElement(&Style{Width: 100}, nil, nil)
		testElement.setWidth(100)

		assert(t, testElement.Bounding.Width == 100, "expected result: %d, but get %d", 100, testElement.Bounding.Width)
	})

	t.Run("set int width 50", func(t *testing.T) {
		testElement := NewHOMElement(&Style{Width: 50}, nil, nil)
		testElement.setWidth(50)

		assert(t, testElement.Bounding.Width == 50, "expected result: %d, but get %d", 100, testElement.Bounding.Width)
	})

	t.Run("set percentage width", func(t *testing.T) {
		t.Run("width 100%", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Width: "100%"}, nil, nil)
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width == 100, "expected result: %d, but get %d", 100, testElement.Bounding.Width)
		})

		t.Run("width 75%", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Width: "75%"}, nil, nil)
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width == 75, "expected result: %d, but get %d", 100, testElement.Bounding.Width)
		})

		t.Run("width 25%", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Width: "25%"}, nil, nil)
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width == 25, "expected result: %d, but get %d", 100, testElement.Bounding.Width)
		})

		t.Run("broken string", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Width: "fuck you"}, nil, nil)
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width == 0, "expected result: %d, but get %d", 100, testElement.Bounding.Width)
		})

		t.Run("struct", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Width: Style{}}, nil, nil)
			testElement.setWidth(100)

			assert(t, testElement.Bounding.Width == 0, "expected result: %d, but get %d", 100, testElement.Bounding.Width)
		})

	})
}

func TestElementSetHeight(t *testing.T) {
	t.Run("set int Height 100", func(t *testing.T) {
		testElement := NewHOMElement(&Style{Height: 100}, nil, nil)
		testElement.setHeight(100)

		assert(t, testElement.Bounding.Height == 100, "expected result: %d, but get %d", 100, testElement.Bounding.Height)
	})

	t.Run("set int Height 50", func(t *testing.T) {
		testElement := NewHOMElement(&Style{Height: 50}, nil, nil)
		testElement.setHeight(50)

		assert(t, testElement.Bounding.Height == 50, "expected result: %d, but get %d", 100, testElement.Bounding.Height)
	})

	t.Run("set percentage Height", func(t *testing.T) {
		t.Run("Height 100%", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Height: "100%"}, nil, nil)
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height == 100, "expected result: %d, but get %d", 100, testElement.Bounding.Height)
		})

		t.Run("Height 75%", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Height: "75%"}, nil, nil)
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height == 75, "expected result: %d, but get %d", 100, testElement.Bounding.Height)
		})

		t.Run("Height 25%", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Height: "25%"}, nil, nil)
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height == 25, "expected result: %d, but get %d", 100, testElement.Bounding.Height)
		})

		t.Run("broken string", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Height: "fuck you"}, nil, nil)
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height == 0, "expected result: %d, but get %d", 100, testElement.Bounding.Height)
		})

		t.Run("struct", func(t *testing.T) {
			testElement := NewHOMElement(&Style{Height: Style{}}, nil, nil)
			testElement.setHeight(100)

			assert(t, testElement.Bounding.Height == 0, "expected result: %d, but get %d", 100, testElement.Bounding.Height)
		})

	})
}
