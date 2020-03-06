package bot

// A WeatherResponse struct to map the Entire Response
type WeatherResponse struct {
    Count int `json:"cnt"`
    WeatherList []WeatherList `json:"list"`
    City City `json:"city"`
}

// A WeatherList Struct to map every datetime weather prediction to.
type WeatherList struct {
    DateTime int  `json:"dt"`
    Weather []CurrentWeather  `json:"weather"`
}

// A CurrentWeather Struct to map every weatherID to.
type CurrentWeather struct {
    WeatherId int `json:"id"`
}

// A struct to map our City name annd Country to.
type City struct {
    Name string `json:"name"`
    Country string `json:"country"`
}
