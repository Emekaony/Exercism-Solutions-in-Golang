package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	var message string = fmt.Sprintf("This is the number %.1f", f)
	return message
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	message := fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
	return message
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	// if of type FancyNumber: extract value and convert from string to int
	// return that inv value after conversion. Ignore error edge case.
	switch fnb.(type) {
	case FancyNumber:
		i, _ := strconv.Atoi(fnb.Value())
		return i
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type) {
	case FancyNumber:
		num := fmt.Sprintf("%.1f", float64(ExtractFancyNumber(fnb)))
		return fmt.Sprintf("This is a fancy box containing the number %s", num)
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	value := ""
	// big boy function
	switch i.(type) {
	case int:
		floatValue := float64(i.(int))
		value = DescribeNumber(floatValue)
	case float64:
		// this is called type assertions
		value = DescribeNumber(i.(float64))
	case NumberBox:
		value = DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		value = DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		value = "Return to sender"
	}
	return value
}
