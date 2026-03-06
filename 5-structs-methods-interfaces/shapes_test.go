package structsMethodsInterfaces

import (
	"math"
	"testing"
)

func TestShapes(t *testing.T) {
	rectangle := Rectangle{5.26, 9.20}
	got := Perimeter(rectangle)
	want := 28.92
	compareFloats(t, rectangle, got, want)
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 5, Length: 4.21}, hasArea: 21.05},
		{name: "Circle", shape: Circle{Radius: 5}, hasArea: 78.54},
		{name: "Triangle", shape: Triangle{Base: 5.5, Height: 3.2}, hasArea: 8.8},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			compareFloats(t, tt.shape, tt.shape.Area(), tt.hasArea)
		})

	}
}

func compareFloats(t testing.TB, shape Shape, got, want float64) {
	t.Helper()
	if math.Abs(got-want) > 1e-3 {
		t.Errorf("%#v has area %g want %g", shape, got, want)
	}
}
