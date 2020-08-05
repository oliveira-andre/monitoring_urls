package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	monitoringTimes = 3
	delay           = 5
)

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
	urls := getUrls()

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
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("something goes wrong", err)
	}

	fmt.Println("Testing the site number", index, ":", url)

	if resp.StatusCode == 200 {
		fmt.Println("The website:", url, "is up and running with success!")
	} else {
		fmt.Println("The website:", url, "is with problems")
	}

}

func getUrls() []string {
	var urls []string
	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("something goes wrong", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		trimLine := strings.TrimSpace(line)
		urls = append(urls, trimLine)
	}

	file.Close()

	return urls
}
