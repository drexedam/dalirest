package rest

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

// GetLightPoints is a get endpoint for querying all light point objects
func GetLightPoints(w http.ResponseWriter, r *http.Request) {
	emptyLPOs := make([]LightPointObject, 0)

	if err := json.NewEncoder(w).Encode(emptyLPOs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetLightPointInfo returns information about a specific light point object
func GetLightPointInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Println(params["id"])

	http.Error(w, "Not implemented", http.StatusInternalServerError)
}

// SceneHandler forwards the request based on the request method
func SceneHandler(w http.ResponseWriter, r *http.Request) {
	// TODO Maybe there is a better way
	if r.Method == "GET" {
		getScenes(w, r)
	} else if r.Method == "POST" {
		activateScene(w, r)
	} else {
		http.NotFound(w, r)
	}
}

// getScenes is a get endpoint for querying all available scenes
func getScenes(w http.ResponseWriter, r *http.Request) {
	emptyScenes := make([]GlobalScene, 0)

	if err := json.NewEncoder(w).Encode(emptyScenes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// activateScene is a post endpoint to activate a specific scene
func activateScene(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusInternalServerError)
}