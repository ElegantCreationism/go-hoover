package room

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoom_Create_Room_5x5_dimensions(t *testing.T) {
	// arrange
	dimensions := Dimensions{5, 5}
	testRoom := [][]Coordinate{{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}},
		{{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}},
		{{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}},
		{{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}},
		{{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}},
	}
	// act
	room, err := CreateRoom(dimensions)

	// assert
	fmt.Println(room)
	assert.Nil(t, err)
	assert.Equal(t, testRoom, room)
}

func TestRoom_Create_Room_5x3_dimensions(t *testing.T) {
	// arrange
	dimensions := Dimensions{5, 3}
	testRoom := [][]Coordinate{{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{3, 0}, {3, 1}, {3, 2}},
		{{4, 0}, {4, 1}, {4, 2}},
	}
	// act
	room, err := CreateRoom(dimensions)

	// assert
	fmt.Println(room)
	assert.Nil(t, err)
	assert.Equal(t, testRoom, room)
}
