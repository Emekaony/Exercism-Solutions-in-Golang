package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// two numbers must be of the same type
	return float64(productionRate) * float64(successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// 10 cars cost $95,000
	// each car cost $10,000
	var groupsOfTen int = carsCount / 10
	var remaining int = carsCount % 10
	return uint(groupsOfTen*95000 + remaining*10000)
}
