package main

import (
	"fmt"
	"math/rand"
)



type egg struct {
	broken bool
}

func main() {
	var rand rand.Rand
	var highestDropPoint, dropCount int
	var floorMax int = 100
	var floorToCheck int = 1
	var floorFound bool
	dropPoint := rand.Intn(floorMax)
	// bldg has 100 floors
	// one floor is highest can be dropped without breaking
	// egg dropped above floor egg.broken := true
	// egg dropped below floor egg.broken := false
	highestDropPoint, floorFound, dropCount = findEggDropPointBrute(dropPoint, floorToCheck, floorMax, dropCount)
	binarySearchDropPoint(dropPoint, floorFound, floorMax, floorToCheck, dropCount)
	fmt.Println("The highest floor is " + string(highestDropPoint))
}



func findEggDropPointBrute(dropPoint int, floorToCheck int, floorMax int, dropCount int) (highestDropPoint int, floorFound bool, totalDrops int) {
	for i := floorToCheck ; i <= floorMax ; i ++ {
		dropCount++
		if floorToCheck >= dropPoint {
			floorMax = i - 1
			floorFound = true
		}
	}
	return floorMax, floorFound, dropCount
}


func binarySearchDropPoint(dropPoint int, floorFound bool, floorMax int, floorToCheck int, dropCount int) (highestDropPoint int, totalDrops int) {
	for !floorFound {
		mid := floorMax/2
		if mid >= dropPoint {
			highestDropPoint, floorFound, dropCount = findEggDropPointBrute(dropPoint, floorToCheck, floorMax, dropCount)
		}
		if mid < dropPoint {
			floorToCheck = mid
			binarySearchDropPoint(dropPoint, floorFound, floorMax, floorToCheck, dropCount)
		}
		
	}
	return highestDropPoint, dropCount
}
