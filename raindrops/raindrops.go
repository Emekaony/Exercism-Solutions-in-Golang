package raindrops

import (
	"fmt"
	"strings"
)

func Convert(number int) string {
	// builder is a fast way manipulate strings in a mutable way
	builder := strings.Builder{}
	if number%3 == 0 {
		builder.WriteString("Pling")
	}
	if number%5 == 0 {
		builder.WriteString("Plang")
	}
	if number%7 == 0 {
		builder.WriteString("Plong")
	}
	if number%7 != 0 && number%3 != 0 && number%5 != 0 {
		// does this work??
		builder.WriteString(fmt.Sprint(number))
	}
	return builder.String()
}
