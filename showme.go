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
const idQueryURL string = "show-details?q="

type titleQueryResponse struct {
	Total    string
	Page     int
	Pages    int
	Tv_shows []map[string]interface{}
}

type idQueryResponse struct {
	TvShow map[string]interface{}
}

func main() {
	var url = createURL()
	// fmt.Println("url:", url)

	var id = getShowID(url)
	// fmt.Println("id:", id)

	var episodes = getEpisodesByID(id)
	fmt.Println("episodes:", episodes)
	// getEpisodesByID(id)
}

func createURL() string {
	flag.Parse()

	if flag.Arg(0) == "" {
		fmt.Println("Please provide a TV show title to search!")
		os.Exit(0)
	}

	return formatTitleQueryURL(flag.Arg(0))
}

func formatTitleQueryURL(title string) string {
	// TODO: replace with regex?
	var formattedTitle = strings.Replace(title, " ", "%20", -1)
	formattedTitle = strings.Replace(formattedTitle, "_", "%20", -1)

	url := strings.Join([]string{apiBaseURL, titleQueryURL, formattedTitle}, "")
	fmt.Println("url:", url)
	return url
}

func getShowID(url string) int {
	var jsonResp = queryShowTitle(url)

	totalTitles, err := strconv.Atoi(jsonResp.Total)
	if err != nil {
		fmt.Println("Error converting jsonResp.Total to integer:", err)
		os.Exit(0)
	}

	switch {
	case totalTitles < 1:
		fmt.Println("No results found with that title! Please try again.")
	case totalTitles > 1:
		fmt.Println("More than one result found! Please narrow down and try again.")
		// TODO: print out list of found TV show titles in console
	default:
		foundTitle := jsonResp.Tv_shows[0]["name"]
		foundID := jsonResp.Tv_shows[0]["id"]
		fmt.Println("foundTitle:", foundTitle)
		fmt.Println("foundID", foundID)
	}

	return int(jsonResp.Tv_shows[0]["id"].(float64))
}

func queryShowTitle(url string) titleQueryResponse {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in Get request:", err)
		os.Exit(0)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading queryShowTitle Body:", err)
		os.Exit(0)
	}

	var jsonResp titleQueryResponse
	json.Unmarshal([]byte(body), &jsonResp)
	// fmt.Println("jsonResp:", jsonResp)

	return jsonResp
}

func formatIDQueryURL(id int) string {
	url := strings.Join([]string{apiBaseURL, idQueryURL, strconv.Itoa(id)}, "")
	// fmt.Println("url:", url)
	return url
}

func getEpisodesByID(id int) []interface{} {
	var url = formatIDQueryURL(id)
	// fmt.Println("url:", url)

	var jsonResp = queryShowID(url)
	// fmt.Println(`jsonResp.TvShow["episodes"]`, jsonResp.TvShow["episodes"])
	return jsonResp.TvShow["episodes"].([]interface{})
}

func queryShowID(url string) idQueryResponse {
	fmt.Println("url:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in Get request:", err)
		os.Exit(0)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading queryShowID Body:", err)
	}

	var jsonResp idQueryResponse
	json.Unmarshal([]byte(body), &jsonResp)
	// fmt.Println("jsonResp:", jsonResp)

	return jsonResp
}
