package purchase
import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car":
		return true
	case "truck":
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	optionSlice := []string{option1, option2}
	sort.Strings(optionSlice)
	answer := fmt.Sprintf("%s is clearly the better choice.", optionSlice[0])
	return answer

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	percentage := 0
	if age < 3 {
		percentage = 80
	} else if age >= 10 {
		percentage = 50
	} else {
		percentage = 70
	}
	calculePrice := float64(originalPrice) * float64(percentage) / 100
	return calculePrice
}