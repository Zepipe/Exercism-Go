package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
    if !ok {
        return false
    }
	_, ok = bill[item]
    if !ok {
        bill[item] = units[unit]
    } else {
        bill[item] += units[unit]
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
    if !ok {
        return false // if the given unit is not in the units map.
    }

    _, ok = bill[item]
    if !ok {
        return false // if the given item is not in the bill
    } 

    res := bill[item] - units[unit]
    
    switch {
        case res < 0:
        	return false
    	case res == 0:
        	delete(bill, item)
        	return true
        default:
        	bill[item] = res
        	return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemVal, ok := bill[item]
    if !ok {
        return 0, false
    } else {
        return itemVal, true 
    }
}
