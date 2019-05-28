package room

import (
	"errors"
	"log"
	"strings"
)

var (
	RoombaPosition Coordinate
)

type Coordinate struct {
	X int
	Y int
}

type Dimensions struct {
	rows    int
	columns int
}

type Patches []Coordinate

func CreateRoom(dimensions Dimensions) [][]Coordinate{
	var room = make([][]Coordinate, dimensions.rows)
	for i := range room {
		room[i] = make([]Coordinate, dimensions.columns)
	}

	for i := 0; i < dimensions.rows; i++ {
		for j := 0; j < dimensions.columns; j++ {
			c := NewCoordinate(i, j)
			room[i][j] = c
		}
	}

	return room
}

func Navigate(instructions string, roombaPosition Coordinate, dimensions Dimensions, patches Patches) (Coordinate, int, error) {
	var patchesCleaned int
	Patches := patches
	isValidInput := strings.Contains(instructions, ",")
	if !isValidInput {
		return roombaPosition, patchesCleaned, errors.New("Invalid Input: Please use comma separated values:\ni.e: `N,S,E,W`")
	}

	room:= CreateRoom(dimensions)
	isValidMove := checkIsValidMove(dimensions, room)
	if !isValidMove {
		return roombaPosition, patchesCleaned, errors.New("Invalid movement")
	}

	s := strings.Split(instructions, ",")
	for i := range s {
		switch true {
		//North
		case "N" == s[i]:
			//Go up
			roombaPosition = goNorth(roombaPosition)

			//Check that you have not moved beyond the bounds of the room
			isValidMove := checkIsValidMove(dimensions, room)
			if !isValidMove {
				return roombaPosition, patchesCleaned, errors.New("Invalid movement")
			}

			// Check if you have landed on a dirt patch and increase the counter
			for patch := range Patches{
				if roombaPosition == Patches[patch]{
					patchesCleaned++
				}
			}
		// East
		case "E" == s[i]:
			// Go right
			roombaPosition = goEast(roombaPosition)

			//Check that you have not moved beyond the bounds of the room
			isValidMove := checkIsValidMove(dimensions, room)
			if !isValidMove {
				return roombaPosition, patchesCleaned, errors.New("Invalid movement")
			}

			// Check if you have landed on a dirt patch and increase the counter
			for patch := range Patches {
				if roombaPosition == Patches[patch]{
					patchesCleaned++
				}
			}

		// South
		case "S" == s[i]:
			// Go left
			roombaPosition = goSouth(roombaPosition)

			isValidMove := checkIsValidMove(dimensions, room)
			if !isValidMove {
				return roombaPosition, patchesCleaned, errors.New("Invalid movement")
			}

			// Check if you have landed on a dirt patch and increase the counter
			for patch := range Patches{
				if roombaPosition == Patches[patch]{
					patchesCleaned++
				}
			}

		// West
		case "W" == s[i]:
			//Go West
			roombaPosition = goWest(roombaPosition)

			isValidMove := checkIsValidMove(dimensions, room)
			if !isValidMove {
				return roombaPosition, patchesCleaned, errors.New("Invalid movement")
			}
			// Check if you have landed on a dirt patch and increase the counter
			for patch := range Patches{
				if roombaPosition == Patches[patch]{
					patchesCleaned++
				}
			}

		default:
			return roombaPosition, patchesCleaned, errors.New("Invalid Input: Please use valid chacters:\nNorth: N\nEast: E\nWest: W\nSouth: S")
		}
	}

	log.Printf("The roomba has Cleaned %v patches", patchesCleaned)
	return roombaPosition, patchesCleaned, nil
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{X: x, Y: y}
}

func NewDimensions(rows int, columns int) Dimensions {
	return Dimensions{rows: rows, columns: columns}
}

func createDirtPatches(coordinates []Coordinate) Patches {
	patches := make(Patches, len(coordinates))
	for  coord := range coordinates {
		patches = append(patches, coordinates[coord])
	}
	return patches
}

func goNorth(coordinate Coordinate) Coordinate{
	return Coordinate{coordinate.X - 1, coordinate.Y}
}

func goEast(coordinate Coordinate) Coordinate{
	return Coordinate{coordinate.X, coordinate.Y + 1}
}

func goSouth(coordinate Coordinate) Coordinate{
	return Coordinate{coordinate.X + 1, coordinate.Y}
}

func goWest(coordinate Coordinate) Coordinate{
	return Coordinate{coordinate.X, coordinate.Y - 1}
}

func checkIsValidMove(dimensions Dimensions, room [][]Coordinate) bool {
	var validMove bool
	for i := 0; i < dimensions.rows; i++ {
		for j := 0; j < dimensions.columns; j++ {
			if RoombaPosition == room[i][j]{
				validMove = true
				break
			}
		}
		// To stop searching once we have verified the coordinates lie in the room
		if validMove {
			break
		}
	}
	// return false if move is invalid
	return validMove
}
