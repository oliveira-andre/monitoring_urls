package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
			getLogs()
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

func chooseOption() (choosedOption int) {
	fmt.Scan(&choosedOption)

	return
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

func getUrls() []string {
	var urls []string
	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("something goes wrong", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		trimLine := strings.TrimSpace(line)
		urls = append(urls, trimLine)
	}

	return urls
}

func checkUrl(index int, url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("something goes wrong", err)
	}

	fmt.Println("Testing the site number", index, ":", url)

	if resp.StatusCode == 200 {
		setLog(url, true)
		fmt.Println("The website:", url, "is up and running with success!")
	} else {
		setLog(url, false)
		fmt.Println("The website:", url, "is with problems")
	}

}

func setLog(url string, status bool) {
	file, err := os.OpenFile("tmp/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("something goes wrong", err)
	}

	defer file.Close()

	formatedCurrentTime := time.Now().Format("02/01/2006 15:04:05")
	strStatus := strconv.FormatBool(status)
	logMessage := formatedCurrentTime + " - " + url + " - online: " + strStatus + "\n"
	file.WriteString(logMessage)
}

func getLogs() {
	fmt.Println("showing logs")

	file, err := ioutil.ReadFile("tmp/log.txt")

	if err != nil {
		fmt.Println("something goes wrong", err)
	}

	fmt.Println(string(file))
}
