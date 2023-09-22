package structures

type FSAPIRawData_DSS struct {
	Server   DSS_ServerStruct   `json:"server"`
	Slots    DSS_SlotsStruct    `json:"slots"`
	Vehicles []DSS_VehicleArray `json:"vehicles"`
	Mods     []DSS_ModArray     `json:"mods"`
	Fields   []DSS_FieldArray   `json:"fields"`
}

type DSS_ServerStruct struct {
	DayTime             int    `json:"dayTime"`
	Game                string `json:"game"`
	MapName             string `json:"mapName"`
	MapSize             string `json:"mapSize"`
	MapOverviewFilename string `json:"mapOverviewFilename"`
	Money               int    `json:"money"`
	Name                string `json:"name"`
	Server              string `json:"server"`
	Version             string `json:"version"`
}

type DSS_SlotsStruct struct {
	Capacity int               `json:"capacity"`
	Used     int               `json:"used"`
	Players  []DSS_PlayerArray `json:"players"`
}

type DSS_PlayerArray struct {
	IsUsed  bool    `json:"isUsed"`
	IsAdmin bool    `json:"isAdmin"`
	Uptime  int     `json:"uptime"`
	Name    string  `json:"name"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Z       float64 `json:"z"`
}

type DSS_VehicleArray struct {
	Name       string          `json:"name"`
	Category   string          `json:"category"`
	Type       string          `json:"type"`
	X          float64         `json:"x"`
	Y          float64         `json:"y"`
	Z          float64         `json:"z"`
	Fills      []DSS_FillArray `json:"fills,omitempty"`
	Controller string          `json:"controller,omitempty"`
}

type DSS_FillArray struct {
	Type  string  `json:"type"`
	Level float64 `json:"level"`
}

type DSS_ModArray struct {
	Author      string `json:"author"`
	Hash        string `json:"hash"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type DSS_FieldArray struct {
	ID      int     `json:"id"`
	IsOwned bool    `json:"isOwned"`
	X       float64 `json:"x"`
	Z       float64 `json:"z"`
}
