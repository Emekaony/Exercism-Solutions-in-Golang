package elon

import (
	"fmt"
)

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.batteryDrain < c.battery {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	message := fmt.Sprintf("Driven %d meters", c.distance)
	return message
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	message := fmt.Sprintf("Battery at %d%s", c.battery, "%")
	return message
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	// this is the amount of times the car can move
	numTimes := int(c.battery / c.batteryDrain)
	// this is the max distance the car can move
	maxDistance := c.speed * numTimes
	if maxDistance >= trackDistance {
		return true
	} else {
		return false
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
