package main

import "fmt"

func main() {

	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am Superman"
	fmt.Println(superman)

	thor := "I am thor" //when datatype is not declared
	fmt.Println(thor)

	// thorRating := 3.
	// fmt.Printf("%v, %T", thorRating, thorRating)

	var Ironman, CapAmerica = "I am Ironman ", "I am CapAmerica"
	fmt.Println(Ironman, CapAmerica)

	var defaultValue string
	fmt.Println(defaultValue) // Empty space  and in string its 0

	var (
		spiderman = "I am spiderman"
		age       = 18
		powers    = "web sling, spidysense"
		antman    = "I am antman"
	)

	fmt.Println(spiderman, age, powers, antman)
}
