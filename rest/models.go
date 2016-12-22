package rest

// TODO Determine which data types fit

// LightPointObject represents a single light point object
type LightPointObject struct {
	Id           string
	Address      string
	Location     string
	Destination  byte
	Intensity    string
	GroupAddress int
	OnOff        bool
	Type         string
}

// GlobalScene represents a global scene
type GlobalScene struct {
	Id     string
	Name   string
	Scenes []InternScene
}

// InternScene represents a intern scene
type InternScene struct {
	Id           string
	Name         string
	LightPoints  []*LightPointSaturation
	sceneId      byte
}

// LightPointSaturation represents a LightPointObject with Saturation
type LightPointSaturation struct {
	LightPoint *LightPointObject
	Saturation  string
}
