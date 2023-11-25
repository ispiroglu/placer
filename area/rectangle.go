package area

type Rectangle struct {
	Id, Width, Height uint8
	p                 Point
	Rotated, Placed   bool
}

func NewRectangle(Id, Width, Height uint8) *Rectangle {
	return &Rectangle{
		Id:      Id,
		Width:   Width,
		Height:  Height,
		Rotated: false,
		p:       Point{X: 0, Y: 0},
	}
}

func (r *Rectangle) Point() *Point {
	return &r.p
}
