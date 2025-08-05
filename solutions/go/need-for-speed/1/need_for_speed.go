package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

func NewTrack(distance int) Track {
	return Track{distance: distance}
}

func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
	return car
}

func CanFinish(car Car, track Track) bool {
	enough := (car.battery / car.batteryDrain) * car.speed
	if enough < track.distance {
		return false
	}
	return true
}
