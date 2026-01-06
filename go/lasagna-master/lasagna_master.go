package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	pasta := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			pasta += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return pasta, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scaledQuantities int) []float64 {
	result := make([]float64, len(quantities))
	for i, quantity := range quantities {
		result[i] = quantity * float64(scaledQuantities) / 2
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
