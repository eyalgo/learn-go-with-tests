package structs

import "testing"

func _assertCorrectMessage(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	_assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	_checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %.g", got, want)
		}
	}
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		_checkArea(t, rectangle, 72.0)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}
		_checkArea(t, circle, 314.1592653589793)
	})
}
