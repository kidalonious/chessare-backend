package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{"message": "Hello from the chessare-backend"}
	json.NewEncoder(w).Encode(response)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	inputQuery := r.URL.Query().Get("q")
	games, err := GetGamesByPlayer(inputQuery)
	if err != nil {
		fmt.Println("it bwoke :(")
		return
	}
	var gamesAsMap []map[string]string
	for _, game := range games {
		temp := GetGameMap(game)
		gamesAsMap = append(gamesAsMap, temp)
	}

	
	response := map[string][]map[string]string{inputQuery: gamesAsMap}  // this needs to be the map of the username to the string map of the games
	json.NewEncoder(w).Encode(response)
}