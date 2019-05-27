package room

type Coordinate struct {
	x int
	y int
}

type Dimensions struct {
	x int
	y int
}


func newCoordinate(x int, y int) Coordinate{
	return Coordinate{x:x, y: y}
}


func CreateRoom(dimensions Dimensions) (interface{}, error) {
	rows := dimensions.x
	columns := dimensions.y
	var room = make([][]Coordinate, rows)
	for i := range room {
		room[i] = make([]Coordinate, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			c := newCoordinate(i,j)
			room[i][j] = c
		}
	}
	return room, nil
}
