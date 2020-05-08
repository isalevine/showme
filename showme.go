package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	const apiBaseURL string = "https://www.episodate.com/api/"
	const titleQueryURL string = "search?q="

	flag.Parse()

	title := flag.Arg(0)
	url := strings.Join([]string{apiBaseURL, titleQueryURL, title}, "")
	fmt.Println("url:", url)
}
