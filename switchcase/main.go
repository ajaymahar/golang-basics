package main

import "fmt"

func main() {
	var favSports string
	fmt.Print("Enter Your Fav sports: ")
	fmt.Scanf("%s", &favSports)

	switch favSports {
	case "cricket":
		fmt.Println("Nice I like ", favSports, "too :)")
	case "football":
		fmt.Println("Nice I play ", favSports, "too :)")
		fallthrough
	case "race":
		fmt.Println("Nice I watch ", favSports, "too :)")
	default:
		fmt.Println("Will do play ", favSports)
	}
}
