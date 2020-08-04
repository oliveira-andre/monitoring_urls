package main

import "fmt"
import "os"
import "net/http"
import "time"

const monitoringTimes = 3
const delay = 5

func main() {
	welcome()

	for {
		showOptions()
		choosedOption := chooseOption()

		switch choosedOption {
		case 0:
			fmt.Println("exiting")
			os.Exit(0)
		case 1:
			monitoring()
		case 2:
			fmt.Println("showing logs")
		default:
			fmt.Println("not valid value choosed")
			os.Exit(-1)
		}
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
	urls := []string{"https://alura.com.br/",
		"https://caelum.com.br"}

	for i := 0; i < monitoringTimes; i++ {
		for index, url := range urls {
			checkUrl(index, url)
		}

		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}

	fmt.Println("")
}

func checkUrl(index int, url string) {
	resp, _ := http.Get(url)
	fmt.Println("Testing the site number", index, ":", url)

	if resp.StatusCode == 200 {
		fmt.Println("The website:", url, "is up and running with success!")
	} else {
		fmt.Println("The website:", url, "is with problems")
	}

}
