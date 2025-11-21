package shapes

import "testing"

func TestArea(t *testing.T) {
	t.Run("Rectangle Area", func(t *testing.T) {
		r := Rectangle{height: 10, width: 15}
		receivedArea := r.Area()
		expectedArea := float64(150.0)

		if receivedArea != expectedArea {
			t.Error("Area received is different from expected")
		}
	})

	t.Run("Circle Area", func(t *testing.T) {
		c := Circle{radius: 10}
		receivedArea := c.Area()
		expectedArea := float64(314)

		if receivedArea != expectedArea {
			t.Error("Area received is different from expected")
		}
	})
}