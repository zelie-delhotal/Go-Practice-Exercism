// Package weather provides weather forecast tools for Goblinocus.
package weather

var (
    // CurrentCondition represents the current weather at a certain place.
	CurrentCondition string
    // CurrentLocation represents a city in Goblinocus.
	CurrentLocation  string
)

// Forecast returns a string descripting the current weather condition at a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
