package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

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
	showTitle, id := getShowTitleAndID("https://www.episodate.com/api/search?q=30%%20Rock")
	if showTitle != expectedOutputShowTitle {
		t.Errorf("getShowTitleAndID(\"https://www.episodate.com/api/search?q=30%%20Rock\") failed, expected %v, got %v", expectedOutputShowTitle, showTitle)
	}
	if id != expectedOutputShowID {
		t.Errorf("getShowTitleAndID(\"https://www.episodate.com/api/search?q=30%%20Rock\") failed, expected %v, got %v", expectedOutputShowID, id)
	}
}

func TestGetEpisodesByID(t *testing.T) {
	episodes := getEpisodesByID(11020)
	if len(episodes) != 138 {
		t.Errorf("getEpisodesByID(11020) failed, expected length to be 138, got %v", len(episodes))
	}
}

func TestSelectRandomEpisode(t *testing.T) {
	// TODO: To decrease likeliness of accidentally selecting same 2 episodes,
	// can also just use `episodes := getEpisodesByID(11020)`
	// (but then this becomes an *integration* test, not a *unit* test)
	episode1 := map[string]interface{}{
		"name":    "Pilot",
		"season":  1.0,
		"episode": 1.0,
	}
	episode2 := map[string]interface{}{
		"name":    "The Aftermath",
		"season":  1.0,
		"episode": 2.0,
	}
	episode3 := map[string]interface{}{
		"name":    "Blind Date",
		"season":  1.0,
		"episode": 3.0,
	}
	episode4 := map[string]interface{}{
		"name":    "Jack the Writer",
		"season":  1.0,
		"episode": 4.0,
	}
	episodes := make([]interface{}, 0)
	episodes = append(episodes, episode1, episode2, episode3, episode4)

	randomEpisode1 := selectRandomEpisode(episodes)
	time.Sleep(1 * time.Second)
	randomEpisode2 := selectRandomEpisode(episodes)
	if randomEpisode1["name"] == randomEpisode2["name"] {
		t.Errorf("selectRandomEpisode(episodes) failed, expected '%v' and '%v' to be different episode names", randomEpisode1["name"], randomEpisode2["name"])
	}
}

func TestFormatEpisodeTitle(t *testing.T) {
	episode := map[string]interface{}{
		"name":    "Pilot",
		"season":  1.0,
		"episode": 1.0,
	}
	episodeTitle := formatEpisodeTitle(episode)
	if !strings.Contains(episodeTitle, episode["name"].(string)) {
		t.Errorf("formatEpisodeTitle(episode) failed, expected %v to contain %v", episodeTitle, episode["name"])
	}
}
