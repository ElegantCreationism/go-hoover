package roomba

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
