// Package weather is made for providing forecasts in different cities in the Goblinocus.
package weather

var (
    // CurrentCondition is used to show weather conditiong of the given city.
	CurrentCondition string
	// CurrentLocation is used to show the given city name for forecast.
    CurrentLocation  string
)
// Forecast is used to show the given city name and current weather in this city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
