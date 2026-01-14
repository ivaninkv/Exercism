package space

type Planet string

const secondsInEarthYear = 31_557_600

func Age(seconds float64, planet Planet) float64 {
	period := getOrbPeriod(planet)
	if period > 0 {
		return seconds / (secondsInEarthYear * period)
	}
	return -1.0
}

func getOrbPeriod(planet Planet) float64 {
	period := 0.0
	switch planet {
	case "Mercury":
		period = 0.2408467
	case "Venus":
		period = 0.61519726
	case "Earth":
		period = 1.0
	case "Mars":
		period = 1.8808158
	case "Jupiter":
		period = 11.862615
	case "Saturn":
		period = 29.447498
	case "Uranus":
		period = 84.016846
	case "Neptune":
		period = 164.79132
	default:
		period = -1.0
	}
	return period
}
