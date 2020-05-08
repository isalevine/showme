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

	// TODO: replace with regex?
	var title = strings.Replace(flag.Arg(0), " ", "%20", -1)
	title = strings.Replace(title, "_", "%20", -1)

	url := strings.Join([]string{apiBaseURL, titleQueryURL, title}, "")
	fmt.Println("url:", url)
}
