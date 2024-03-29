package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#vgot %g want %g", tt.shape, got, tt.want)
		}
	}
}

// checkArea := func(t testing.TB, shape Shape, want float64) {
// 	t.Helper()
// 	got := shape.Area()
// 	if got != want {
// 		t.Errorf("got %f want %f", got, want)
// 	}
// }
// t.Run("rectangles", func(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
// 	want := 72.0
// 	checkArea(t, rectangle, want)
// })

// t.Run("circles", func(t *testing.T) {
// 	circle := Circle{10}
// 	want := 314.1592653589793
// 	checkArea(t, circle, want)
// })
