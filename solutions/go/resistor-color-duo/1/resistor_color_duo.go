package resistorcolorduo

import ("unicode/utf8"
        "strconv"
       )

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) == 0 {
        return 0
    }
    
	tmpValue := "" // String representation of return int value
	bandColors := map[string]int {
        "black" : 0, "brown" : 1, "red" : 2,
        "orange" : 3, "yellow" : 4, "green" : 5,
        "blue" : 6, "violet" : 7, "grey" : 8, "white" : 9,
    }

    for _, color := range colors {
        colorValue, exists := bandColors[color]
        if exists {
            // check if black band is first color
            if color == "black" && tmpValue == "" {
                continue
            }
            tmpValue += strconv.Itoa(colorValue)
            // check if value has 2 color bands already
            if utf8.RuneCountInString(tmpValue) == 2 {
                break
            }
        }
    }

    if tmpValue == "" {
        return 0
    }

    res, _ := strconv.Atoi(tmpValue)
    return res
}
