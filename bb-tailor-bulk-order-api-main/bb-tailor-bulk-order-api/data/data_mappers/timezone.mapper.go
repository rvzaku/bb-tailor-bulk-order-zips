package data_mappers

type Timezone struct {
	ZoneName      string `json:"zoneName"`
	GmtOffset     int    `json:"gmtOffset"`
	GmtOffsetName string `json:"gmtOffsetName"`
	Abbreviation  string `json:"abbreviation"`
	TzName        string `json:"tzName"`
}
