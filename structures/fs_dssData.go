package structures

type FSAPIRawDataResponse struct {
	Server   FSAPI_ServerObj    `json:"server"`
	Slots    FSAPI_SlotsObj     `json:"slots"`
	Vehicles []FSAPI_VehicleArr `json:"vehicles"`
	Mods     []FSAPI_ModArr     `json:"mods"`
	Fields   []FSAPI_FieldArr   `json:"fields"`
}

type FSAPI_ServerObj struct {
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

type FSAPI_SlotsObj struct {
	Capacity int               `json:"capacity"`
	Used     int               `json:"used"`
	Players  []FSAPI_PlayerArr `json:"players"`
}

type FSAPI_PlayerArr struct {
	IsUsed  bool    `json:"isUsed"`
	IsAdmin bool    `json:"isAdmin"`
	Uptime  int     `json:"uptime"`
	Name    string  `json:"name"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Z       float64 `json:"z"`
}

type FSAPI_VehicleArr struct {
	Name     string          `json:"name"`
	Category string          `json:"category"`
	Type     string          `json:"type"`
	X        float64         `json:"x"`
	Y        float64         `json:"y"`
	Z        float64         `json:"z"`
	Fills    []FSAPI_FillArr `json:"fills"`
}

type FSAPI_FillArr struct {
	Type  string  `json:"type"`
	Level float64 `json:"level"`
}

type FSAPI_ModArr struct {
	Author      string `json:"author"`
	Hash        string `json:"hash"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type FSAPI_FieldArr struct {
	ID      int     `json:"id"`
	IsOwned bool    `json:"isOwned"`
	X       float64 `json:"x"`
	Z       float64 `json:"z"`
}
