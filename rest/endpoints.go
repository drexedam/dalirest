package rest

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strings"
	"github.com/drexedam/daliclient"
	"fmt"
)

var (
	lpos []*LightPointObject
	scenes []*InternScene
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

	scenes = make([]*InternScene, 6)
	scenes[0] = &InternScene{
		Id: "1",
		Name: "Scene 1",
		LightPoints: []*LightPointSaturation{
			{
				LightPoint: lpos[0],
				Saturation: "10e-2",
			},
			{
				LightPoint: lpos[1],
				Saturation: "0",
			},
			{
				LightPoint: lpos[2],
				Saturation: "0",
			},
		},
		sceneId: 0x00,
	}

	scenes[1] = &InternScene{
		Id: "2",
		Name: "Scene 2",
		LightPoints: []*LightPointSaturation{
			{
				LightPoint: lpos[0],
				Saturation: "51e-2",
			},
			{
				LightPoint: lpos[1],
				Saturation: "100e-2",
			},
			{
				LightPoint: lpos[2],
				Saturation: "51e-2",
			},
		},
		sceneId: 0x01,
	}

	scenes[2] = &InternScene{
		Id: "3",
		Name: "Scene 3",
		LightPoints: []*LightPointSaturation{
			{
				LightPoint: lpos[0],
				Saturation: "100e-2",
			},
			{
				LightPoint: lpos[1],
				Saturation: "49e-2",
			},
			{
				LightPoint: lpos[2],
				Saturation: "51e-2",
			},
		},
		sceneId: 0x02,
	}


	scenes[3] = &InternScene{
		Id: "4",
		Name: "Scene 4",
		LightPoints: []*LightPointSaturation{
			{
				LightPoint: lpos[1],
				Saturation: "19e-2",
			},
			{
				LightPoint: lpos[2],
				Saturation: "19e-2",
			},
		},
		sceneId: 0x03,
	}

	scenes[4] = &InternScene{
		Id: "5",
		Name: "Scene 5",
		LightPoints: []*LightPointSaturation{
			{
				LightPoint: lpos[1],
				Saturation: "7e-2",
			},
			{
				LightPoint: lpos[2],
				Saturation: "7e-2",
			},
		},
		sceneId: 0x04,
	}

	scenes[5] = &InternScene{
		Id: "6",
		Name: "Scene 6",
		LightPoints: []*LightPointSaturation{
			{
				LightPoint: lpos[1],
				Saturation: "3e-2",
			},
			{
				LightPoint: lpos[2],
				Saturation: "3e-2",
			},
		},
		sceneId: 0x04,
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

// getScenes is a get endpoint for querying all available scenes
func GetScenes(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(scenes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// activateScene is a post endpoint to activate a specific scene
func ActivateScene(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, scene := range scenes {
		if strings.Compare(scene.Id, id) == 0 {
			usbdali := &daliclient.Usbdali{}

			if err := usbdali.Connect("localhost"); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer usbdali.Close()

			if err := usbdali.Send(daliclient.MakeBroadcastCmd(scene.sceneId, daliclient.CmdSetScene)); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var resp []byte
			var err error
			if resp, err = usbdali.Receive(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Println(resp)
		}
	}
}