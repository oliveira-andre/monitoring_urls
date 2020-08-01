package main

import "fmt"

func main() {
  name := "andre"
  version := 1.1

  fmt.Println("Hello mr.", name)
  fmt.Println("this project is on version", version)

  fmt.Println("1 - initialize monitoring")
  fmt.Println("2 - show logs")
  fmt.Println("0 - exit")

  var choosed_option int
  fmt.Scan(&choosed_option)
}
