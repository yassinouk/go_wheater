package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		/* Lat     float32 `json:"lat"` */
	} `josn:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			/* Icon string `json:"icon"` */
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
					/* Icon string `json:"icon"` */
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	/* colorReset := "\033[0m" */

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	/* this is to reset the color if you want to*/
	/* fmt.Println("test", string(colorReset)) */
	/* url := "http://api.weatherapi.com/v1/" */
	/* params := [5]string{"current.json", "forecast.json", "marine.json", "stronomy.json", "ip.json"} */

	apiKey := "d30b83540f874ee085e123807233007"
	city := "Rabat"

	resp, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + apiKey + "&q=" + city + "&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("something went wrong! wheater api is not reachable for today")
	}
	body, err := ioutil.ReadAll(resp.Body)
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf("%s, %s: %.0f°C, %s\n",
		location.Name, location.Country, current.TempC, current.Condition.Text,
	)
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue

		}
		message := fmt.Sprintf("%s|---> %.0f°C, %.0f%%, %s\n",
			date.Format("15:01"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		switch {
		case hour.ChanceOfRain <= 40:
			fmt.Print(string(colorGreen), message)
		case hour.ChanceOfRain <= 60:
			fmt.Print(string(colorYellow), message)
		case hour.ChanceOfRain > 60:
			fmt.Print(string(colorRed), message)
		}
	}
}
