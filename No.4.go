package main

import "fmt"

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
	n1Len := len(nums1)
	n2Len := len(nums2)
	half := (n1Len + n2Len) / 2

	n1left := n1Len/2 - 1
	n1right := n1Len / 2
	n2left := n2Len/2 - 1
	n2right := n2Len / 2
	fmt.Print(half)

	for 
	// if n1Len%2 == 0 {
	// 	n1Middle := n1Len / 2
	// }
	// if n2Len%2 == 0 {
	// 	n2Middle := n2Len / 2
	// }
	return 0
}

func main() {
	x := 1
	fmt.Print(x)
}
