package data_mappers

type Country struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Iso3           string     `json:"iso3"`
	Iso2           string     `json:"iso2"`
	NumericCode    string     `json:"numeric_code"`
	PhoneCode      string     `json:"phone_code"`
	Capital        string     `json:"capital"`
	Currency       string     `json:"currency"`
	CurrencyName   string     `json:"currency_name"`
	CurrencySymbol string     `json:"currency_symbol"`
	Tld            string     `json:"tld"`
	Native         string     `json:"native"`
	Region         string     `json:"region"`
	RegionID       string     `json:"region_id"`
	Subregion      string     `json:"subregion"`
	SubregionID    string     `json:"subregion_id"`
	Nationality    string     `json:"nationality"`
	Timezones      []Timezone `json:"timezones"`
	Latitude       string     `json:"latitude"`
	Longitude      string     `json:"longitude"`
	States         []State    `json:"states"`
}
