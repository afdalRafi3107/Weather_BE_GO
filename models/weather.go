package models

type WeatherData struct {
	Name string `json:"name"`

	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"temp"`
	} `json:"main"`

	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:icon`
	} `json:"weather"`

	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}