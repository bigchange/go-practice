/*
 * @Author: Jerry You
 * @CreatedDate: 2023/3/15
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/3/15 15:13
 */

package exercism

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (r TemperatureUnit) String() string {
	uints := []string {"°C", "°F"}
	return uints[r]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (r Temperature) String() string {
	return fmt.Sprintf("%v %v", r.degree, r.unit)
}
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (r SpeedUnit) String() string {
	uints := []string{"km/h", "mph"}
	return uints[r]
}

func (r Speed) String() string {
	return fmt.Sprintf("%v %v", r.magnitude, r.unit)
}
// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
// <location>: <temperature>, Wind <wind_direction> at <wind_speed>, <humidity>% Humidity
func (r MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", r.location, r.temperature, r.windDirection, r.windSpeed, r.humidity)
}
