# What is showme?
showme is a simple Go script that takes a TV show title as an argument, and queries the Episodate API to find a random episode title for you to watch!


# Setup
## Requirements
* Go version 1.14.2 or higher
* Written and tested in macOS Mojave version 10.14.6

## Installation
* Clone this repo. 
* Navigate to the cloned `/showme` directory.
* You can now run the script with `go run showme.go "tv show title"`!

If you want to build the binary:
* Run `go build showme.go`.
* You can now run the binary with `./showme "tv show title"`!


# Use
Without building the binary, you can run `go run showme.go "tv show title"`.
After building the binary, you can run `./showme "tv show title"`.


## Examples
Before building binary:
```
$ go run showme.go "30 rock"

// Output:

OK! From the show '30 Rock', you should watch:

Season 4, Episode 22 - I Do Do

Enjoy!
```

After building binary:
```
$ ./showme "30 rock"

// Output:

OK! From the show '30 Rock', you should watch:

Season 4, Episode 22 - I Do Do

Enjoy!
```

# Notes
## Planned Additions:
* When multiple TV show titles are found, print a list of titles to the console to help users re-search for the correct show.
* Add a flag to set the number of episodes returned.
* Add option to pass title argument via `title=tv_show_name` (code already handles parsing underscores in title)
* Add functionality to create aliases for quicker searching, i.e. alias `koth` for `King of the Hill`
