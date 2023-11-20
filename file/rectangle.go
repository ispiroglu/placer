package file

type Rectangle struct {	
	Id, X, Y, Width, Height uint8
	Rotated, Placed             bool 
}

func NewRectangle(Id, Width, Height uint8) *Rectangle {
	return &Rectangle{
		Id: Id,
		Width: Width,
		Height: Height,
		Rotated: false,
		X: 0,
		Y: 0,
	}
}