package lasagna

func PreparationTime(layers []string, avgPrepTime int) int {
    if len(layers) == 0 {
        return 0
    }

    if avgPrepTime == 0 {
    	return len(layers) * 2 
    } else {
        return len(layers) * avgPrepTime
    }
}

func Quantities(layers []string) (int, float64) {
    var noodlesWeight int
    var litreSauce float64
    
    for _, val := range layers {
        switch val {
            case "noodles":
                noodlesWeight += 50
            case "sauce":
                litreSauce += 0.2
            default:
                continue
        }
    }
    
    return noodlesWeight, litreSauce
}

func AddSecretIngredient(friendsList, ownList []string) {
    if len(friendsList) != 0 && len(ownList) != 0 {
        ownList[len(ownList) - 1] = friendsList[len(friendsList) - 1]
    }
}

func ScaleRecipe(amounts []float64, numOfPortions int) []float64 {
    var scaledQuants []float64
	var scaledQuant	float64

    if numOfPortions <= 0 || len(amounts) == 0 {
        return amounts
    }
    
    for idx, _ := range amounts {
        scaledQuant = amounts[idx] / 2.0 * float64(numOfPortions) 
    	scaledQuants = append(scaledQuants, scaledQuant)
    }
    
    return scaledQuants
}

