package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function

func DivideFood(f FodderCalculator, numCows int) (float64, error) {
	totalFodders, err1 := f.FodderAmount(numCows)
	if err1 != nil {
		return 0.0, err1
	}
	// will multiply it by this factor below
	factor, err2 := f.FatteningFactor()
	if err2 != nil {
		return 0.0, err2
	}

	return (totalFodders * factor) / float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function

func ValidateInputAndDivideFood(f FodderCalculator, numCows int) (float64, error) {
	if numCows > 0 {
		return DivideFood(f, numCows)
	}
	return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsErr struct {
	numberOfCows int
	message      string
}

// this is cool
func (i *InvalidCowsErr) Error() string {
	// learned that u should always use pointer references for
	// Error implementation. Will understand why later on.
	message := fmt.Sprintf("%d cows are invalid: %s", i.numberOfCows, i.message)
	return message
}

// TODO: define the 'ValidateNumberOfCows' function

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		// here we pass a reference as return value
		return &InvalidCowsErr{
			numberOfCows: numCows,
			message:      "there are no negative cows",
		}
	} else if numCows == 0 {
		return &InvalidCowsErr{
			numberOfCows: numCows,
			message:      "no cows don't need food",
		}
	} else {
		return nil
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
