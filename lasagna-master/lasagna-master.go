package lasagna

// PreparationTime calculates preparation time
func PreparationTime(layers []string, t int) int {
	if t == 0 {
		t = 2
	}
	return len(layers) * t
}

// Quantities calculates quantities of noodles and sauce
func Quantities(layers []string) (int, float64) {
	var noodles int = 0
	var sauce float64 = 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		default:
			continue
		}
	}
	return noodles, sauce
}

// AddSecretIngredient adds secret ingridient to my list
func AddSecretIngredient(frList []string, myList []string) []string {
	return append(myList, frList[len(frList)-1])
}

// ScaleRecipe scales recipe for given number of portions
func ScaleRecipe(quants []float64, s int) []float64 {
	scaled := []float64{}
	for _, twoPortions := range quants {
		scaled = append(scaled, (twoPortions/2.0)*float64(s))
	}
	return scaled
}
