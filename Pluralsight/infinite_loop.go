package main

func main() { //value of i is implicitly declared
	var i int // declare i as a int data type
	for ; ;  // no intializer, no conditional check, and no post clause
		if i == 5 {
			break //loop is stopped by the i == 5 conditional. Remove at own risk!
		}
		println(i)
		i++
	}
}