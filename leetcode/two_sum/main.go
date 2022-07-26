package main

func main() {

}

func twoSumBrute(nums []int, target int) []int{
	// return indicies of two numbers 
	// signed or unsigned
	// such that they add up to target
	// gaurenteed one solution
	var solution []int

	for i := 0 ; i < len(nums) ; i ++ {
		// if nums[i] <= target { doesn't work for unsigned
		for j := i+1 ; j < len(nums) ; j++ {
			if nums[i] + nums[j] == target { //o (n^2)
				solution = append(solution, i)
				solution = append(solution, j)
				return solution
			}
		}
		// }
	}
	return solution
}


func twoSumHashMap(nums []int, target int) []int{
	// return indicies of two numbers 
	// signed or unsigned
	// such that they add up to target
	// gaurenteed one solution
	// var solution []int
	m := make(map[int]int)
	for i := 0 ; i < len(nums) ; i ++ { // o(n)
		elem, ok := m[target - nums[i]]
		if ok {
			// solution = append(solution, i)
			// solution = append(solution, elem)
			return []int{elem, i}
		}
		m[nums[i]] = i  // memory complexity o(n)
	}
	return nil
}

//consider constraints?
