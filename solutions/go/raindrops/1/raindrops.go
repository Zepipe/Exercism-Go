package raindrops

import "strconv"

func Convert(number int) string {
    var resStr string

    if number % 3 == 0 {
        resStr += "Pling"
    }

    if number % 5 == 0 {
        resStr += "Plang"
    }
    
	if number % 7 == 0 {
        resStr += "Plong"
    } 
    
    if resStr == "" {
        return strconv.Itoa(number)
    } else {
        return resStr
    }
}
