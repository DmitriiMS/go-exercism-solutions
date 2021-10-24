package space

type Planet string

const EarthYearInSeconds = 31557600

var planetYears = map[Planet]float64{
	"Mercury": EarthYearInSeconds * 0.2408467,
	"Venus":   EarthYearInSeconds * 0.61519726,
	"Earth":   EarthYearInSeconds,
	"Mars":    EarthYearInSeconds * 1.8808158,
	"Jupiter": EarthYearInSeconds * 11.862615,
	"Saturn":  EarthYearInSeconds * 29.447498,
	"Uranus":  EarthYearInSeconds * 84.016846,
	"Neptune": EarthYearInSeconds * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	return float64(seconds) / planetYears[planet]
}
