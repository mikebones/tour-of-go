package main

func main() {

}


func search(nums []int, target int) int {
	// nums is sorted in ascending order
	// if target exists return its index
	left := 0
	right := len(nums)-1
	indexToSearch := left + (right - left) / 2
	for left <= right {
		switch {
		case target == nums[indexToSearch] :
			return indexToSearch
		case target > nums[indexToSearch] :
			// nums = nums[indexToSearch-1:]
			// indexToSearch = len(nums)/2
			left = indexToSearch + 1
			indexToSearch = (left + right)/2
		case target < nums[indexToSearch] :
			// nums = nums[:indexToSearch]
			// indexToSearch = len(nums)/2
			right = indexToSearch - 1
			indexToSearch = (left + right)/2
		case indexToSearch == len(nums)-1 || indexToSearch == 0 :
			return -1
		}
	}	
	// nums[:val/2]
	return -1

}

// halfway point
// larger or smaller
// 

func findval(nums []int, valToSearch int) {

}

// slice operator is hanf-open range
// includes first but exludes last
