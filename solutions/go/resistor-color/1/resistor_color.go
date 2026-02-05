package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string {"black", "brown", "red", "orange", "yellow",
	"green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colors := Colors()
	colorResistance := 0
    
    for idx, val := range colors {
    	if color == val {
            colorResistance = idx
        }
    }

    return colorResistance
}
