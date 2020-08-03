package main

import "fmt"
import "os"
import "net/http"

func main() {
	welcome()
	showOptions()
	choosedOption := chooseOption()

	switch choosedOption {
	case 0:
		fmt.Println("exiting")
	case 1:
		monitoring()
	case 2:
		fmt.Println("showing logs")
		os.Exit(0)
	default:
		fmt.Println("not valid value choosed")
		os.Exit(-1)
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

func monitoring() {
	fmt.Println("initializing the monitoring...")
	url := "https://www.alura.com.br"
	resp, _ := http.Get(url)

	if resp.StatusCode == 200 {
		fmt.Println("The website:", url, "is up and running with success!")
	} else {
		fmt.Println("The website:", url, "is with problems")
	}
}
