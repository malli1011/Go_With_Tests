package stuctsinterfaces

import (
	"fmt"
	"math"
	"testing"
)

func TestPrintAreas(t *testing.T) {

	var shapes []Shape

	shapes = append(shapes, Circle{5.0})
	shapes = append(shapes, Square{5.0})
	shapes = append(shapes, Rectangel{5.0, 10.0})

	for _, shape := range shapes {
		fmt.Printf("Aread of %T is %f \n", shape, shape.getArea())
	}
}

func TestArea(t *testing.T) {
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{2.5}
		got := circle.getArea()
		expected := Area(math.Pi * math.Pow(2.5, 2))
		if expected != got {
			t.Errorf("got %g expected %g", got, expected)
		}
	})
	t.Run("Squear", func(t *testing.T) {
		square := Square{5}
		got := square.getArea()
		expected := Area(25)

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}

	})
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangel{5, 10}
		got := rect.getArea()
		expected := Area(50)

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}

	})
}

func TestArea2(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangel{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.getArea()
		if float64(got) != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
