package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	res := float64(productionRate) * successRate  / 100
    return res
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	res := int(float64(productionRate) * successRate)  / 100 /60
    return res
}

func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	individualCars := carsCount % 10
	res := groupsOfTen*95000 + individualCars*10000
	return uint(res)
}
