package contains_duplicate

func main() {
	
}

func ContainsDuplicate(nums []int) (ContainsDuplicate bool) {
	var compare []int
	for i := 0 ; i < len(nums) ; i++ {			// o(n)
		for c := 0 ; c < len(compare) ; c++ {   // o(n)
			if nums[i] == compare[c] {
				return true
			}
		}
		compare = append(compare, nums[i])
	}
	return false
}



// TODO: review this result

// func containsDuplicate(nums []int) bool {
//     set := make(map[int]struct{})
//     for _, v := range nums {
//         if _, ok := set[v]; ok {
//             return true
//         } else {
//             set[v] = struct{}{}
//         }
//     }
//     return false
// }
