package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ElegantCreationism/go-hoover/room"
	"github.com/ElegantCreationism/go-hoover/roomba"
	"io/ioutil"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "HOME: %v\n", "This is a Roomba simulation application service, send a POST request to http://0.0.0.0:8080/roomba ")
}

func RoombaHandler(w http.ResponseWriter, r *http.Request) {
	//Create data stucture to unmarshal JSON body from request to the roomba
	roombaRequest := roomba.NewRequestBody()
	// Read the Body and handle the Error
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("failed to read body of request: %v", err.Error())
		return
	}
	// Unmarshal body into the Data structure
	err = json.Unmarshal(body, &roombaRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("failed to unmarshal body of request: handler.go: line 54")
		return
	}
	// Separate the body into variable so we can create the room and navigate to clean dirt patches
	instructions, startCoords, dimensions, patches := roomba.SeparateVariablesInRequestBody(roombaRequest)
	// Pass variables to Navigate which will create the room using the dimensions
	// and Navigate
	endpoint, dirtPatchesCleaned, err := room.Navigate(instructions, startCoords, dimensions, patches)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("INVALID INPUT: %v", err.Error())
		return
	}
	endpointArray := []int{endpoint.X, endpoint.Y}

	roombaResponse := roomba.NewResponseBody(endpointArray, dirtPatchesCleaned)

	data, err := json.Marshal(roombaResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("failed to marshal data structure into a response: handler.go: line 52")
		return
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
