package main

import "fmt"

func main() {
  welcome()
  showOptions()
  choosedOption := chooseOption()

  switch choosed_option {
    case 0:
      fmt.Println("exiting")
    case 1:
      fmt.Println("initializing the monitoring...")
    case 2:
      fmt.Println("showing logs")
    default:
      fmt.Println("not valid value choosed")
  }
}

func welcome() {
  name := "andre"
  version := 1.1

  fmt.Println("Hello mr.", name)
  fmt.Println("this project is on version", version)
}

func showOptions() {
  fmt.Println("1 - initialize monitoring")
  fmt.Println("2 - show logs")
  fmt.Println("0 - exit")
}

func chooseOption() int {
  var choosedOption int
  fmt.Scan(&choosedOption)

  return choosedOption
}
