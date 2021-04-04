package main
import (
	"fmt"
)
func main() {
	helloList := []string{
		'Yoroshiku',
		'剣道',
		'Done it',
		'nyet',
		'Aegon Targaryen',
	}
	fmt.Println(len(helloList))
	fmt.Println(helloList[len(helloList)-1])
	fmt.Println(helloList[len(helloList)])
}
