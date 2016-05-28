package main

import "fmt"

func main() {
  fmt.Println("What would you like to convert?")
  fmt.Println("1: Binary")
  fmt.Println("2: Hex")
  var user_input int
  var converted int

  fmt.Scanln(&user_input)

    switch user_input {
  case 1: fmt.Println("Binary conversion, Enter your number: ")
          fmt.Scanln(&converted)
          fmt.Printf("Your converted binary number is %b\n" , converted)
  case 2: fmt.Printf("%x", &user_input)
  default: fmt.Println("not a known input")

  }
}
