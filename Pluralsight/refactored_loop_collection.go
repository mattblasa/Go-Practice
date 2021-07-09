package main

func main() {
	slice := []int{1,2,3} //declare array
	for i:=0; i<len(slice); i++ { //start from 0, if i is less than the slice, increment 1
		println(slice[i]) //print the number in the array
	}
}