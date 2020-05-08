package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in Get request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading Body:", err)
	}

	type JsonResponse struct {
		Total    string
		Page     int
		Pages    int
		Tv_shows []map[string]interface{}
	}

	var jsonResp JsonResponse
	json.Unmarshal([]byte(body), &jsonResp)
	fmt.Println("jsonResp:", jsonResp)
}
