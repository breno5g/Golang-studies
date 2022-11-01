package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("result %.2f, expected %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()

		if result != expected {
			t.Errorf("result %.2f, expected %.2f", result, expected)
		}
	}

	t.Run("Rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		expected := 100.0

		verifyArea(t, rectangle, expected)
	})

	t.Run("Circle area", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		verifyArea(t, circle, expected)
	})
}
