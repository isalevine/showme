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

const apiBaseURL string = "https://www.episodate.com/api/"
const titleQueryURL string = "search?q="

func main() {
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("Please provide a TV show title to search!")
		os.Exit(0)
	}

	var title = formatShowTitle(flag.Arg(0))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in Get request:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading Body:", err)
	}

	type JSONResponse struct {
		Total   string
		Page    int
		Pages   int
		TvShows []map[string]interface{}
	}

	var jsonResp JSONResponse
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
		foundTitle := jsonResp.TvShows[0]["name"]
		foundID := jsonResp.TvShows[0]["id"]
		fmt.Println("foundTitle:", foundTitle)
		fmt.Println("foundID", foundID)
	}
}

func formatShowTitle(title string) string {
	// TODO: replace with regex?
	var formattedTitle = strings.Replace(title, " ", "%20", -1)
	formattedTitle = strings.Replace(title, "_", "%20", -1)

	url := strings.Join([]string{apiBaseURL, titleQueryURL, formattedTitle}, "")
	// fmt.Println("url:", url)
	return url
}

func getShowID(title string) int {

}

func getEpisodesByID(id int) []string {

}
