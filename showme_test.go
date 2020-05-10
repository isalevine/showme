package main

import (
	"os"
	"testing"
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
