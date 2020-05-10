package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	url := createURL()
	showTitle, id := getShowTitleAndID(url)
	episodes := getEpisodesByID(id)
	episode := selectRandomEpisode(episodes)
	episodeTitle := formatEpisodeTitle(episode)
	output := strings.Join([]string{"OK! From the show '", showTitle, "', you should watch:\n\n", episodeTitle, "\n\nEnjoy!"}, "")
	fmt.Println(output)
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
	formattedTitle := strings.Replace(title, " ", "%20", -1)
	formattedTitle = strings.Replace(formattedTitle, "_", "%20", -1)
	url := strings.Join([]string{apiBaseURL, titleQueryURL, formattedTitle}, "")
	return url
}

func getShowTitleAndID(url string) (string, int) {
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
	}

	return jsonResp.Tv_shows[0]["name"].(string), int(jsonResp.Tv_shows[0]["id"].(float64))
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
	return jsonResp
}

func formatIDQueryURL(id int) string {
	url := strings.Join([]string{apiBaseURL, idQueryURL, strconv.Itoa(id)}, "")
	return url
}

func getEpisodesByID(id int) []interface{} {
	var url = formatIDQueryURL(id)
	var jsonResp = queryShowID(url)
	return jsonResp.TvShow["episodes"].([]interface{})
}

func queryShowID(url string) idQueryResponse {
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
	return jsonResp
}

func selectRandomEpisode(episodes []interface{}) map[string]interface{} {
	// https://stackoverflow.com/a/33994791
	seed := rand.NewSource(time.Now().Unix())
	randomInt := rand.New(seed)
	episodeNumber := randomInt.Intn(len(episodes))
	episode := episodes[episodeNumber : episodeNumber+1][0].(map[string]interface{})
	return episode
}

func formatEpisodeTitle(episode map[string]interface{}) string {
	formattedTitle := strings.Join([]string{"Season ", strconv.Itoa(int(episode["season"].(float64))), ", Episode ", strconv.Itoa(int(episode["episode"].(float64))), " - ", episode["name"].(string)}, "")
	return formattedTitle
}
