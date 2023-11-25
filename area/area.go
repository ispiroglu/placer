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

func (area *Area) canPlace(rectangle *Rectangle, point *Point) bool {
	// Check if the rectangle fits within the area
	if point.X+rectangle.Width > area.Width || point.Y+rectangle.Height > area.Height {
		return false
	}

	// Check if the space is already occupied
	// TODO: Memoize the occupied spaces faster jumps.
	for i := point.X; i < point.X+rectangle.Width; i++ {
		for j := point.Y; j < point.Y+rectangle.Height; j++ {
			if area.Occupied[i][j] {
				return false
			}
		}
	}

	return true
}

// Should I move this two function to area.go?
func (a *Area) BottomLeftFill() {
	rand.New(rand.NewSource(time.Now().UnixNano()	))

	rand.Shuffle(len(a.Rectangles), func(i, j int) { a.Rectangles[i], a.Rectangles[j] = a.Rectangles[j], a.Rectangles[i] })

    for _, rectangle := range a.Rectangles {
        // Find the bottom-left-most position where the rectangle fits
        point := FindBottomLeftPosition(rectangle, a)
        // Place the rectangle and update the environment
        a.place(rectangle, point)
    }
}

func FindBottomLeftPosition(rect *Rectangle, a *Area) *Point {
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

func (area *Area) place(rectangle *Rectangle, point *Point) {
	// Mark the space as occupied
	for i := point.X; i < point.X+rectangle.Width; i++ {
		for j := point.Y; j < point.Y+rectangle.Height; j++ {
			area.Occupied[i][j] = true
		}
	}
	// Update the rectangle's position
	rectangle.p = *point
	rectangle.Placed = true
}
