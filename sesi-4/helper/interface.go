package helper

type ShapeRepo interface {
	Luas() float64
	Keliling() float64
}

func NewRectange(width, height float64) ShapeRepo {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func NewCircle(radius float64) ShapeRepo {
	return &Circle{
		radius: radius,
	}
}

func (r *Rectangle) Luas() float64 {
	return r.width * r.height
}

func (c *Circle) Luas() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Keliling() float64 {
	return 2 * 3.14 * c.radius
}

func (r *Rectangle) Keliling() float64 {
	return 2 * (r.width + r.height)
}
