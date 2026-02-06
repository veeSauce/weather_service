package page_models

type ForecastModel struct {
	Context  []interface{} `json:"@context"`
	Type     string        `json:"type"`
	Geometry struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"geometry"`
	Properties struct {
		Units             string `json:"units"`
		ForecastGenerator string `json:"forecastGenerator"`
		GeneratedAt       string `json:"generatedAt"`
		UpdateTime        string `json:"updateTime"`
		ValidTimes        string `json:"validTimes"`
		Elevation         struct {
			UnitCode string  `json:"unitCode"`
			Value    float64 `json:"value"`
		} `json:"elevation"`
		Periods []struct {
			Number                     int         `json:"number"`
			Name                       string      `json:"name"`
			StartTime                  string      `json:"startTime"`
			EndTime                    string      `json:"endTime"`
			IsDaytime                  bool        `json:"isDaytime"`
			Temperature                int         `json:"temperature"`
			TemperatureUnit            string      `json:"temperatureUnit"`
			TemperatureTrend           interface{} `json:"temperatureTrend"`
			ProbabilityOfPrecipitation struct {
				UnitCode string `json:"unitCode"`
				Value    int    `json:"value"`
			} `json:"probabilityOfPrecipitation"`
			WindSpeed        string `json:"windSpeed"`
			WindDirection    string `json:"windDirection"`
			Icon             string `json:"icon"`
			ShortForecast    string `json:"shortForecast"`
			DetailedForecast string `json:"detailedForecast"`
		} `json:"periods"`
	} `json:"properties"`
}
