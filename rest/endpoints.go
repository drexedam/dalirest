package rest

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strings"
)

var (
	lpos []*LightPointObject
)

func init() {
	lpos = make([]*LightPointObject, 3)
	lpos[0] = &LightPointObject{
		Id: "1",
		Destination: 0x0,
		Address: "Keks1",
		Location: "U325-Team3",
		Type: "Lamp",
	}
	lpos[1] = &LightPointObject{
		Id: "2",
		Destination: 0x1,
		Address: "Keks2",
		Location: "U325-Team3",
		Type: "Lamp",
	}
	lpos[2] = &LightPointObject{
		Id: "3",
		Destination: 0x2,
		Address: "Keks3",
		Location: "U325-Team3",
		Type: "Lamp",
	}
}

// GetLightPoints is a get endpoint for querying all light point objects
func GetLightPoints(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(lpos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetLightPointInfo returns information about a specific light point object
func GetLightPointInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)



	id := params["id"]
	for _, lpo := range lpos {
		if strings.Compare(lpo.Id, id) == 0 {
			if err := json.NewEncoder(w).Encode(lpo); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.NotFound(w, r)
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