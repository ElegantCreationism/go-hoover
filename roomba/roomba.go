package roomba

import "github.com/ElegantCreationism/go-hoover/room"

type RequestBody struct {
	RoomSize     []int   `json:"roomSize"`
	StartCoords  []int   `json:"coords"`
	Patches      [][]int `json:"patches"`
	Instructions string  `json:"instructions"`
}

type ResponseBody struct {
	EndCoords []int `json:"coords"`
	Patches   int   `json:"patches"`
}

func NewResponseBody(endCoords []int, patches int) ResponseBody {
	return ResponseBody{
		EndCoords: endCoords,
		Patches:   patches,
	}
}

func NewRequestBody() RequestBody {
	return RequestBody{}
}

func SeparateVariablesInRequestBody(requestBody RequestBody) (string, room.Coordinate, room.Dimensions, room.Patches) {
	instructions := requestBody.Instructions
	startCoords := room.NewCoordinate(requestBody.StartCoords[0], requestBody.StartCoords[1])
	dimensions := room.NewDimensions(requestBody.RoomSize[0], requestBody.RoomSize[1])
	patches := make(room.Patches, len(requestBody.Patches))
	for i := 0; i < 1; i++ {
		for j := 0; j < len(requestBody.Patches); j++ {
			x := room.NewCoordinate(requestBody.Patches[0][i], requestBody.Patches[0][i+1])
			patches = append(patches, x)
		}
	}
	return instructions, startCoords, dimensions, patches
}
