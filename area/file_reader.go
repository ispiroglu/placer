package area

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadArea() *Area {
	file, err := os.Open("./Original/C1_1")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rectangleCount uint8 = 0

	if scanner.Scan() {
		c, _ := strconv.ParseUint(scanner.Text(), 10, 8)
		rectangleCount = uint8(c)
	}

	var area *Area
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		width, _ := strconv.ParseUint(parts[0], 10, 8)
		height, _ := strconv.ParseUint(parts[1], 10, 8)

		area = &Area{
			Width:  uint8(width),
			Height: uint8(height),
		}
	}

	rectangles := make([]*Rectangle, rectangleCount)
	var id uint8 = 0
	for scanner.Scan() {
		id += 1
		line := scanner.Text()
		parts := strings.Split(line, " ")
		width, _ := strconv.ParseUint(parts[0], 10, 8)
		height, _ := strconv.ParseUint(parts[1], 10, 8)
		rectangles = append(rectangles, NewRectangle(id, uint8(width), uint8(height)))
	}

	area.Rectangles = rectangles
	area.Occupied = make([][]bool, area.Height)
	for i := range area.Occupied {
		area.Occupied[i] = make([]bool, area.Width)
	}

	return area
}
