package main

import (
	"os"
	"testing"
	"time"
)

// func main() {
// 	flagInput := parseFlag()
// 	url := createTitleQueryURL(flagInput)
// 	showTitle, id := getShowTitleAndID(url)
// 	episodes := getEpisodesByID(id)
// 	episode := selectRandomEpisode(episodes)
// 	episodeTitle := formatEpisodeTitle(episode)
// 	output := strings.Join([]string{"OK! From the show '", showTitle, "', you should watch:\n\n", episodeTitle, "\n\nEnjoy!"}, "")
// 	fmt.Println(output)
// }

func TestParseFlag(t *testing.T) {
	expectedOutput := "30 Rock"
	os.Args = []string{"cmd", "30 Rock"}
	flagInput := parseFlag()
	if flagInput != expectedOutput {
		t.Errorf("parseFlag() failed, expected %v, got %v", expectedOutput, flagInput)
	}
}

func TestCreateTitleQueryURL(t *testing.T) {
	expectedOutput := "https://www.episodate.com/api/search?q=30%20Rock"
	url := createTitleQueryURL("30 Rock")
	if url != expectedOutput {
		t.Errorf("createTitleQueryURL(\"30 Rock\") failed, expected %v, got %v", expectedOutput, url)
	}
}

func TestGetShowTitleAndID(t *testing.T) {
	expectedOutputShowTitle := "30 Rock"
	expectedOutputShowID := 11020
	showTitle, id := getShowTitleAndID("https://www.episodate.com/api/search?q=30%20Rock")
	if showTitle != expectedOutputShowTitle {
		t.Errorf("getShowTitleAndID failed, expected %v, got %v", expectedOutputShowTitle, showTitle)
	}
	if id != expectedOutputShowID {
		t.Errorf("getShowTitleAndID failed, expected %v, got %v", expectedOutputShowID, id)
	}
}

func TestGetEpisodesByID(t *testing.T) {
	episodes := getEpisodesByID(11020)
	if len(episodes) != 138 {
		t.Errorf("getEpisodesByID(11020) failed, expected length to be 138, got %v", len(episodes))
	}
}

func TestSelectRandomEpisode(t *testing.T) {
	episodes := getEpisodesByID(11020)
	episode1 := selectRandomEpisode(episodes)
	time.Sleep(1 * time.Second)
	episode2 := selectRandomEpisode(episodes)
	if episode1["name"] == episode2["name"] {
		t.Errorf("selectRandomEpisode failed, expected '%v' and '%v' to be different episode names", episode1["name"], episode2["name"])
	}
}

func TestFormatEpisodeTitle(t *testing.T) {

}
