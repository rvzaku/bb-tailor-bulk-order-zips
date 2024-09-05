package data_mappers

type State struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StateCode string `json:"state_code"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Type      string `json:"type"`
	Cities    []City `json:"cities"`
}
