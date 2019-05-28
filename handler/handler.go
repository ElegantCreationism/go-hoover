package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ElegantCreationism/go-hoover/room"
	"github.com/ElegantCreationism/go-hoover/roomba"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HOME: %v\n", vars["category"])
}

func RoombaHandler(w http.ResponseWriter, r *http.Request) {
	roombaRequest := roomba.NewRequestBody()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("failed to read body of request: handler.go: line 50")
	}
	err = json.Unmarshal(body, &roombaRequest)
	if err != nil {
		log.Fatal("failed to unmarshal body of request: handler.go: line 54")
	}
	instructions := roombaRequest.Instructions
	startCoords := room.NewCoordinate(roombaRequest.StartCoords[0], roombaRequest.StartCoords[1])
	dimensions := room.NewDimensions(roombaRequest.RoomSize[0], roombaRequest.RoomSize[1])
	patches := make(room.Patches, len(roombaRequest.Patches))
	for i := 0 ; i < 1; i++ {
		for j := 0 ; j < len(roombaRequest.Patches); j++ {
			x := room.NewCoordinate(roombaRequest.Patches[0][i],roombaRequest.Patches[0][i+1])
			patches = append(patches, x)
		}
		//patches = append(patches, room.NewCoordinate(roombaRequest.Patches[0][i], roombaRequest.Patches[0][i+1]))
	}

	endpoint, dirtPatches, err := room.Navigate(instructions, startCoords, dimensions, patches)
	endpointArray := []int{endpoint.X, endpoint.Y}

	roombaResponse := roomba.NewResponseBody(endpointArray, dirtPatches)

	data, err := json.Marshal(roombaResponse)
	if err != nil {
		log.Fatal("failed to unmarshal body of request: handler.go: line 54")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
}
