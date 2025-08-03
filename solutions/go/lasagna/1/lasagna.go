package lasagna

const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actualMinutesInOven`.
func RemainingOvenTime(actualMinutesInOven int) int {
	remaining := OvenTime - actualMinutesInOven
	return remaining
}

// PreparationTime calculates the time needed to prepare the lasagna based on the number of layers.
func PreparationTime(numberOfLayers int) int {
	preparation := numberOfLayers * 2
	return preparation
}

// ElapsedTime calculates the total time spent including preparation and baking time.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	elapsed := PreparationTime(numberOfLayers) + actualMinutesInOven
	return elapsed
}
