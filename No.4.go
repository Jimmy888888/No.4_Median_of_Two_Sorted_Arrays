package main

import (
	"fmt"
	"sort"
)

//pseudo code
/*
total = nums1 + nums2
half = total/2
a : 10   20 30   al ar
b : 1 1  2 4 5   bl br   => al >= br br++, bl >= ar ar++
1 1 2  4 5  10 20 30
target even : half - 1, half
	   odd  : half
target is the end point
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	all := []int{}
	all = append(nums1, nums2...)
	sort.Ints(all)

	var mid1 int = 0
	var mid2 int = 0

	if len(all)%2 == 0 {
		mid1 = len(all) / 2
		mid2 = mid1 - 1
	}
	if len(all)%2 == 1 {
		mid1 = len(all) / 2
		mid2 = mid1
	}
	mid1_num := float64(all[mid1])
	mid2_num := float64(all[mid2])
	var ans float64 = (mid1_num + mid2_num) / 2

	return ans
}

func main() {
	x := 1
	fmt.Print(x)
}
