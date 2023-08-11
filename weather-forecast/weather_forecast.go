// Package weather provides functionalities used to forecast the weather.
package weather

// CurrentCondition represents the current weather condnition as a string.
var CurrentCondition string

// CurrentLocation represents the current location as a string.
var CurrentLocation string

// Forecast function takes the two parameters: city and condition, and returns
// the forecast for that city as a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
