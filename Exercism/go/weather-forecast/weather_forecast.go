// Package weather ...
package weather

// CurrentCondition is the variable to store the current wheather.
var CurrentCondition string

// CurrentLocation is a variable used to store the current city.
var CurrentLocation string

// Forecast retuns the wheater in certain location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
