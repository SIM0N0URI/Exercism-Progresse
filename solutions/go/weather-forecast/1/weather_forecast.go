// Package weather provides weather forecasting tools
// for various cities in Goblinocus.
package weather

// CurrentCondition stores a short description of the current weather condition (e.g., "sunny", "rainy").
var CurrentCondition string

// CurrentLocation holds the name of the city for which the weather is being reported.
var CurrentLocation string

// Forecast returns a string describing the current weather condition at the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
