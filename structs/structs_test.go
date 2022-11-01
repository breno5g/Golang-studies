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
	t.Run("Rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		result := rectangle.Area()
		expected := 100.0

		if result != expected {
			t.Errorf("result %.2f, expected %.2f", result, expected)
		}
	})

	t.Run("Circle area", func(t *testing.T) {
		circle := Circle{10}
		result := circle.Area()
		expected := 314.1592653589793

		if result != expected {
			t.Errorf("result %.2f, expected %.2f", result, expected)
		}
	})
}
