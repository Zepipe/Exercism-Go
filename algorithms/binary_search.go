package binarysearch

import "sort"

func SearchInts(list []int, key int) int {
    // make a copy to save original slice
    sortedList := make([]int, len(list))
    copy(sortedList, list)
    sort.Ints(sortedList)

    low, high := 0, len(list) - 1 // left & right bounds
    for low <= high {
         mid := (low + high) / 2
        if sortedList[mid] == key {
            for mid > 0 && sortedList[mid - 1] == key {
                mid--
            }
            return mid
        }
        // check if mid elem > target elem
        if sortedList[mid] > key {
            high = mid -1
        } else {
            low = mid + 1
        }
    }

    return -1
}  
