package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	const apiBaseURL string = "https://www.episodate.com/api/"
	const titleQueryURL string = "search?q="

	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("Please provide a TV show title to search!")
		os.Exit(0)
	}

	// TODO: replace with regex?
	var title = strings.Replace(flag.Arg(0), " ", "%20", -1)
	title = strings.Replace(title, "_", "%20", -1)

	url := strings.Join([]string{apiBaseURL, titleQueryURL, title}, "")
	// fmt.Println("url:", url)

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
	// fmt.Println("jsonResp:", jsonResp)

	totalTitles, err := strconv.Atoi(jsonResp.Total)
	if err != nil {
		fmt.Println("Error converting jsonResp.Total to integer:", err)
	}

	switch {
	case totalTitles < 1:
		fmt.Println("No results found with that title! Please try again.")
	case totalTitles > 1:
		fmt.Println("More than one result found! Please narrow down and try again.")
		// TODO: print out list of found TV show titles in console
	default:
		foundTitle := jsonResp.Tv_shows[0]["name"]
		foundId := jsonResp.Tv_shows[0]["id"]
		fmt.Println("foundTitle:", foundTitle)
		fmt.Println("foundId", foundId)
	}
}
