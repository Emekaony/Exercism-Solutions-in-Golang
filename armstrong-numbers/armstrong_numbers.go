package armstrong

import "math"

func IsNumber(n int) bool {
	dummyN := n
	digits := []int{}
	cnt := 0
	for dummyN > 0 {
		digits = append(digits, dummyN%10)
		dummyN = dummyN / 10
		cnt++
	}
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum = sum + int(math.Pow(float64(digits[i]), float64(cnt)))
	}
	return sum == n
}
