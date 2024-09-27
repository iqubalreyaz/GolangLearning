package main

import "fmt"

func main() {
	slice1 := []int{3, 8, 9}
	fmt.Printf("\nLength : %d, Capacity: %d & Memory: %p of slice1 before append.", len(slice1), cap(slice1), &slice1)
	slice1 = append(slice1[:], 10)
	slice1 = append(slice1[:], 0)
	slice1 = append(slice1[:], 110)
	slice1 = append(slice1[:], 120)
	//Capacity of Slice doubles itself when you are appending extra element beyond the lentgh of slice
	fmt.Printf("\nLength : %d, Capacity: %d & Memory: %p of slice1 after append.", len(slice1), cap(slice1), &slice1)
}
