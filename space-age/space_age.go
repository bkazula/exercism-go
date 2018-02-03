package space

import (
	"time"
)

type Planet string

var earthYearInSec float64 = 31557600

func Age(sec float64, planet Planet) float64 {
	planetsYears := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	t := time.Second * time.Duration(sec)
	return t.Seconds() / planetsYears[planet] / earthYearInSec
}
