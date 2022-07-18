package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	s := Car{battery: 100, batteryDrain: batteryDrain, speed: speed, distance: 0}
	return s
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{distance: distance}
	return track
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {

	battery := car.battery - car.batteryDrain
	var distance int
	if battery < 0 {
		battery = car.battery
		distance = car.distance
	} else {
		battery = car.battery - car.batteryDrain
		distance = car.speed + car.distance
	}

	return Car{battery: battery, batteryDrain: car.batteryDrain, speed: car.speed, distance: distance}
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

	if (car.battery/car.batteryDrain)*car.speed < track.distance {
		return false
	} else {
		return true
	}

	panic("Please implement the CanFinish function")
}
