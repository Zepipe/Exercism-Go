package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
    // Check if lists are equal
    if areEqual(l1, l2) {
        return RelationEqual
    }
    
    // Check if l1 is a sublist of l2
    if isSublist(l1, l2) {
        return RelationSublist
    }
    
    // Check if l1 is a superlist of l2
    if isSublist(l2, l1) {
        return RelationSuperlist
    }
    
    // If none of the above, they are unequal
    return RelationUnequal
}

// areEqual checks if two slices are equal (same length and same elements in same order)
func areEqual(l1, l2 []int) bool {
    if len(l1) != len(l2) {
        return false
    }
    
    for i := 0; i < len(l1); i++ {
        if l1[i] != l2[i] {
            return false
        }
    }
    
    return true
}

// isSublist checks if sub is a contiguous sub-sequence of super
func isSublist(sub, super []int) bool {
    // An empty list is a sublist of any list
    if len(sub) == 0 {
        return true
    }
    
    // If sub is longer than super, it can't be a sublist
    if len(sub) > len(super) {
        return false
    }
    
    // Check each possible starting position in super
    for i := 0; i <= len(super)-len(sub); i++ {
        found := true
        for j := 0; j < len(sub); j++ {
            if super[i+j] != sub[j] {
                found = false
                break
            }
        }
        if found {
            return true
        }
    }
    
    return false
}