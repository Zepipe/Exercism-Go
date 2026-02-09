package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    
	planetsPeriod := map[Planet]float64 {
        "Mercury" : 0.2408467,
        "Venus" : 0.61519726,
        "Earth" : 1.0,
        "Mars" : 1.8808158,
        "Jupiter" : 11.862615,
        "Saturn" : 29.447498,
        "Uranus" : 84.016846,
        "Neptune" : 164.79132,
    }

    var resAge float64 = -1
    
    period, ok := planetsPeriod[planet]
    
	if ok && seconds != 0 {
        resAge = seconds / period / (365.25 * 24 * 60 * 60) // to get age depends on planet 
    }

    return resAge
}
