package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * avgTime
	}
}

func Quantities(layers []string) (int, float64) {
	sumNoodles := 0
	sumSauce := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sumSauce++
		} else if layers[i] == "noodles" {
			sumNoodles++
		}
	}
	return sumNoodles * 50, sumSauce * 0.2
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portion int) []float64 {
	// precompute
	var tt = make([]float64, len(quantities))
	mulitplier := float64(portion) / 2
	for i := 0; i < len(quantities); i++ {
		tt[i] = quantities[i] * mulitplier
	}
	return tt
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
