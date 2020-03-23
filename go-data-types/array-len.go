package main
import "fmt"

func setArrayValue(array [4]int) {
	array[0] = 2
}

func setSliceValue(slice []int) {
	slice[0] = 2
	slice = append(slice, 5)
}

func main() {
	// anArray := [4]int{1, 0, 0, -4}
	// fmt.Println("Array length:", len(anArray))
	// fmt.Println("Value at index 3:", anArray[3])

	// aSlice := []int{1,2,3,4}
	// fmt.Println(cap(aSlice))

	// setArrayValue(anArray)
	// fmt.Println(anArray)

	// sliceTwo := anArray[:2]
	// fmt.Println(sliceTwo)

	// setSliceValue(aSlice)
	// fmt.Println(aSlice)

	// Creates a nil slice of integers

	// Create a slice of strings with the slice literal syntax:
	aSlice := make([]int, 5)
fmt.Println(aSlice)
}