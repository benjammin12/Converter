package main

import "fmt"

func main() {
	var (
		userInput int //choice for value you'd like to converted to
		converted int //hold number you would like to convert
	)

	fmt.Println("What would you like to convert?") //print out choices
	fmt.Println("1: Binary")
	fmt.Println("2: Hex")
	fmt.Println("3: Unicode")

	fmt.Scanln(&userInput)

	switch userInput {
	case 1: //user picked one, for Binary
		fmt.Println("Binary conversion, Enter your number: ")
		fmt.Scanln(&converted) //store value the user enters
		fmt.Printf("Your converted binary number is %b\n", converted)
	case 2: //user picked two, for Hexi
		fmt.Println("Hexidecimal conversion,  Enter your number: ")
		fmt.Scanln(&converted)
		fmt.Printf("Your converted Hexidecimal number is %X\n", converted)
	case 3: //user picked three, for unicode
		fmt.Println("Unicode Conversion, Enter your number: ")
		fmt.Scanln(&converted)
		fmt.Printf("Your converted Unicode number is %U\n", converted)
	default: //you didn't pick one through thee
		fmt.Println("not a known input")

	}
}
