package area

import (
	"math/rand"
	"time"
)

type Area struct {
	Width, Height uint8
	Rectangles    []*Rectangle
	Occupied      [][]bool
}

type Point struct {
	X, Y uint8
}

func (a *Area) BottomLeftFill() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(a.Rectangles), func(i, j int) { a.Rectangles[i], a.Rectangles[j] = a.Rectangles[j], a.Rectangles[i] })

	for _, rectangle := range a.Rectangles {
		// Find the bottom-left-most position where the rectangle fits
		point := findBottomLeftPosition(rectangle, a)
		// Place the rectangle and update the environment
		a.place(rectangle, point)
	}
}

// ! What if cannout find a position?
func findBottomLeftPosition(rect *Rectangle, a *Area) *Point {
	zero := uint8(0)
	for y := zero; y < a.Height-rect.Height; y++ {
		for x := zero; x < a.Width-rect.Width; x++ {
			point := &Point{X: x, Y: y}
			if a.canPlace(rect, point) {
				return point
			}
		}
	}
	// If no position is found, return the start position
	return &Point{X: zero, Y: zero}
}

func (a *Area) place(rectangle *Rectangle, point *Point) {
	// Mark the space as occupied
	for i := point.X; i < point.X+rectangle.Width; i++ {
		for j := point.Y; j < point.Y+rectangle.Height; j++ {
			a.Occupied[i][j] = true
		}
	}
	// Update the rectangle's position
	rectangle.p = *point
	rectangle.Placed = true
}

func (a *Area) canPlace(rectangle *Rectangle, point *Point) bool {
	// Check if the rectangle fits within the area
	if point.X+rectangle.Width > a.Width || point.Y+rectangle.Height > a.Height {
		return false
	}

	// Check if the space is already occupied
	// TODO: Memoize the occupied spaces faster jumps.
	for i := point.X; i < point.X+rectangle.Width; i++ {
		for j := point.Y; j < point.Y+rectangle.Height; j++ {
			if a.Occupied[i][j] {
				return false
			}
		}
	}

	return true
}
