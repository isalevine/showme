package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const apiBaseURL string = "https://www.episodate.com/api/"
const titleQueryURL string = "search?q="
const idQueryURL string = "show-details?q="

type titleQueryResponse struct {
	Total   string
	Page    int
	Pages   int
	TvShows []map[string]interface{} `json:"tv_shows"`
}

type idQueryResponse struct {
	TvShow map[string]interface{}
}

func main() {
	flagInput := parseFlag()
	url := createTitleQueryURL(flagInput)
	showTitle, id := getShowTitleAndID(url)
	episodes := getEpisodesByID(id)
	episode := selectRandomEpisode(episodes)
	episodeTitle := formatEpisodeTitle(episode)
	output := strings.Join([]string{"OK! From the show '", showTitle, "', you should watch:\n\n", episodeTitle, "\n\nEnjoy!"}, "")
	fmt.Println(output)
}

func parseFlag() string {
	flag.Parse()
	if flag.Arg(0) == "" {
		log.Fatal("Please provide a TV show title to search!")
	}
	return flag.Arg(0)
}

func createTitleQueryURL(flagInput string) string {
	// TODO: check for aliases matching flagInput
	// TODO: replace with regex?
	formattedTitle := strings.Replace(flagInput, " ", "%20", -1)
	formattedTitle = strings.Replace(formattedTitle, "_", "%20", -1)
	url := strings.Join([]string{apiBaseURL, titleQueryURL, formattedTitle}, "")
	return url
}

func getShowTitleAndID(url string) (string, int) {
	var jsonResp = queryShowTitle(url)
	totalTitles, err := strconv.Atoi(jsonResp.Total)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case totalTitles < 1:
		log.Fatal("No results found with that title! Please try again.")
	case totalTitles > 1:
		log.Fatal("More than one result found! Please narrow down and try again.")
		// TODO: print out list of found TV show titles in console
	}

	return jsonResp.TvShows[0]["name"].(string), int(jsonResp.TvShows[0]["id"].(float64))
}

func queryShowTitle(url string) titleQueryResponse {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
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
