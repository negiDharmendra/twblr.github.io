package interfaces

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length int
	breadth int
}

type Hybrid struct {
	square Square
	rectangle Rectangle
}

func (this *Square) Area() int{
	return this.side * this.side
}
func (this *Rectangle) Area() int{
	return this.length * this.breadth
}

func (this *Hybrid) Area() int{
	return this.square.Area() + this.rectangle.Area()
}