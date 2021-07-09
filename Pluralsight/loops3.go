package main

func main() { //value of i is implicitly declared
	for i := 0; i < 5; i++ { // this defines i within the loop. It is not globally declared. 
		println(i)
	}
}