// Package weather is responsible for showing the current condition of a provided city.
package weather

// CurrentCondition represents a current condition.
var CurrentCondition string

// CurrentLocation represents a current location.
var CurrentLocation string

// Forecast returns a forecast based on city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
