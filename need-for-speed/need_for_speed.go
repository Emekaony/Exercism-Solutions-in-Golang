package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
	return car
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	t := Track{
		distance: distance,
	}
	return t
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	} else {
		car.distance += car.speed
		car.battery -= car.batteryDrain
		return car
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	numTimes := int(car.battery / car.batteryDrain)
	// this is because the car covers a distance of speed meters everytime you drive it
	maxDistance := car.speed * numTimes
	if maxDistance >= track.distance {
		return true
	} else {
		return false
	}
}
