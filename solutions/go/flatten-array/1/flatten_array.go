package flatten

func Flatten(nested any) []any {
	result := make([]any, 0)
	
	// Helper function to recursively process the nested structure
	var flattenHelper func(any)
	flattenHelper = func(item any) {
		if item == nil {
			return
		}
		
		// Check if the item is a slice
		if arr, ok := item.([]any); ok {
			// Recursively process each element in the slice
			for _, elem := range arr {
				flattenHelper(elem)
			}
		} else {
			// For non-slice items, add to result
			result = append(result, item)
		}
	}
	
	flattenHelper(nested)
	return result
}