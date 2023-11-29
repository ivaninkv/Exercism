// Package weather_forecast implements a simple weather forecast program.
package weather

// CurrentCondition - current weather condition.
var CurrentCondition string

// CurrentLocation - current location.
var CurrentLocation string

// Forecast - current weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
