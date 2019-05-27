package roomba


type roombaRequestBody struct {
	RoomSize []int `json:"roomSize"`
	StartCoords []int `json:"coords"`
	Patches [][]int `json:"patches"`
	Instructions string `json:"instructions"`
}

type roombaResponseBody struct{
	EndCoords []int `json:"coords"`
	Patches int `json:"patches"`
}

func NewRequestBody() roombaRequestBody {
	return roombaRequestBody{}
}


func NewResponseBody() roombaResponseBody{
	return roombaResponseBody{}
}


