package maketest_test

import "testing"

func printSlice(s []int) {
	// fmt.Printf("len=%d cap=%d %v", len(s), cap(s), s)
	print(len(s))
	print(cap(s))
}

func printArray(a [100]int) {
	// fmt.Printf("len=%d cap=%d %v", len(a), cap(a), a)
	print(len(a))
	print(cap(a))
}

func TestMain(t *testing.T) {
	s := make([]int, 100)
	printSlice(s)

	var a [100]int
	printArray(a)
}
