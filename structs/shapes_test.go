package shapes

import "testing"

var tests = []struct {
	name         string
	shape        Shape
	hasPerimeter float64
	hasArea      float64
}{
	{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerimeter: 36.0, hasArea: 72.0},
	{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586, hasArea: 314.1592653589793},
	{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasPerimeter: 28.97056274847714, hasArea: 36.0},
}

func TestPerimeter(t *testing.T) {
	for _, tt := range tests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
		}
	}
}

func TestArea(t *testing.T) {
	for _, tt := range tests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
