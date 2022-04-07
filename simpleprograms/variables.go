package main

// Have to remember these are () and NOT {}
import (
	"fmt"
	"strconv"
)

func main() {
	var text = "String"
	fmt.Println(text)

	var number int = 7
	// strconv is needed to convert int to string
	fmt.Println(text, strconv.Itoa(number))
	// Could also do this
	fmt.Println(text + " " + strconv.Itoa(number))

	
}
