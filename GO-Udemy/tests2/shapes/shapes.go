package shapes

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
}