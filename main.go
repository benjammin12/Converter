package main

import "fmt"

func main() {
  fmt.Println("What would you like to convert?")
  var user_input string

  fmt.Scanf("%s", &user_input)

    switch user_input {
  case "binary": fmt.Println("Binary conversion, Enter your number: ")
  case "hex": fmt.Printf("%x", &user_input)
  default: fmt.Println("not a known input")

  }
}
