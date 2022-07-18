package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
	panic("NeedsLicense not implemented")
}

var s string

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	for i := 0; i < len(option1); i++ {
		if option1[i] < option2[i] {
			return fmt.Sprintf("%v is clearly the better choice.", option1)
			break
		} else if option1[i] > option2[i] {
			return fmt.Sprintf("%v is clearly the better choice.", option2)
			break
		} else if option1[i] == option2[i] {
		}
	}
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	} else if 3 <= age && age < 10 {
		return originalPrice * 0.7
	} else if age >= 10 {
		return originalPrice * 0.5
	}

	panic("CalculateResellPrice not implemented")
}
