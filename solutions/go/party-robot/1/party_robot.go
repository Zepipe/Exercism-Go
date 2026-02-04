package partyrobot

import "strconv"
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "!" + " You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var tableStr string
    var distanceStr string
    
    if table < 100 && table >= 10 {
        tableStr = "0" + strconv.Itoa(table)
    } else if table < 10 {
        tableStr = "00" + strconv.Itoa(table)
    } else {
        tableStr = strconv.Itoa(table)
    }
	
    distanceStr = fmt.Sprintf("%.1f", distance)
    
    return Welcome(name) + "\n" + "You have been assigned to table " + tableStr + ". " + "Your table is " + direction + ", exactly " + distanceStr + " meters from here.\n" + "You will be sitting next to " + neighbor + "."
}
