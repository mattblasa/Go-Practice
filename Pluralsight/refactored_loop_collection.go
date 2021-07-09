package main

func main() {
	slice := []int{1, 2, 3} //declare array
	for i, v := range slice { // for i is the index in the array, and slice is the the number in the array
		println(i, v) //prints te
	}
}