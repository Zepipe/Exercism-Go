package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if productionRate <= 0 || successRate <= 0 {
        return 0.0
    } else {
        return float64(productionRate) / 100.0 * successRate
    }
	
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	if productionRate <= 0 || int(successRate) <= 0 {
        return 0
    } else {
        return int( float64(productionRate) / 100.0 * successRate / 60.0 )
    }
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var individualCars int
    
	if carsCount <= 0 {
        return 0
    } else if carsCount % 10 == 0 {
        return uint(carsCount * 9500)
    } else {
 		individualCars = carsCount % 10
        return uint((carsCount - individualCars) * 9500 + individualCars * 10000)
    }
}
