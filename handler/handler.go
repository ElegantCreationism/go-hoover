package handler

import (
	"encoding/json"
	"fmt"
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
	roombaRequest := newRequestBody()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("failed to read body of request: handler.go: line 50")
	}
	err = json.Unmarshal(body, &roombaRequest)
	if err != nil {
		log.Fatal("failed to unmashal body of request: handler.go: line 54")
	}






	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


}

