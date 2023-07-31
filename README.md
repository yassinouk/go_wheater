# Weather CLI App in Golang

This CLI app is fetching data from a weather API and displaying it in the terminal with colored output.

## Project Information

- Project: weather-cli
- Created Date: 30/07/2023

## Introduction

This Golang CLI application fetches weather data from a weather API and presents it in the terminal with color-coded output based on the chance of rain. It uses the [Weather API](https://www.weatherapi.com/) to retrieve the current weather conditions and the hourly forecast for a specified city.

## Usage

To use the app, you need to have Golang installed on your machine. If you don't have Golang installed, you can download it from the [official Golang website](https://go.dev/).

1. Clone the repository to your local machine:

```bash
git clone https://github.com/yassinouk/go_wheater.git
```
2. Navigate to the project directory:

```bash
cd go_weather
```
3. In the main.go file change the variable city with your desired city:

4. Build the app:

```bash
go build -o go_weather
```
5. Run the app:

```bash
./go_weather
```
## Example output
```bash
$ go run main.go
London, United Kingdom: 19°C, Light rain
15:07|---> 20°C, 85%, Patchy rain possible
16:07|---> 21°C, 0%, Partly cloudy
17:07|---> 21°C, 0%, Partly cloudy
18:07|---> 20°C, 0%, Partly cloudy
19:07|---> 20°C, 0%, Partly cloudy
20:07|---> 18°C, 0%, Partly cloudy
21:07|---> 17°C, 0%, Partly cloudy
22:07|---> 16°C, 0%, Partly cloudy
23:07|---> 16°C, 0%, Partly cloudy
```