# What is showme?
`showme` is a simple Go script that takes a TV show title as an argument, and queries the Episodate API to find a random episode title for you to watch!

# Highlights
### 1. This is the first Go script I've written!
I learned enough Go from scratch to write this in one weekend.

Even though it obviously doesn't take advantage of Go's concurrency, it's nice to take a break from Ruby and play with static typing! I always learn a lot about how data structures are implemented under-the-hood in dynamic languages. 

In this case, [parsing JSON from an API](https://github.com/isalevine/showme/blob/e43fa4be557503224474214e004805a198106a35/showme.go#L90) and [learning about interface{}](https://github.com/isalevine/showme/blob/e43fa4be557503224474214e004805a198106a35/showme.go#L122) have helped me understand how ambiguous data structures (like an API response) have memory allocated!

### 2. I wrote this with TDD, and [all the functions called in `main()` have unit tests](https://github.com/isalevine/showme/blob/master/showme_test.go)!
Several interesting situations I encountered were:
* [Mocking command-line argument flags](https://github.com/isalevine/showme/blob/8409ba6eb1357f3726817c71f6bf7117ec730a60/showme_test.go#L12)
* [Testing pseudo-random number generator functionality](https://github.com/isalevine/showme/blob/8409ba6eb1357f3726817c71f6bf7117ec730a60/showme_test.go#L74) that uses [`time.Now().Unix()` as the seed for the `rand` object](https://github.com/isalevine/showme/blob/e43fa4be557503224474214e004805a198106a35/showme.go#L124)
* Figuring out the [syntax needed to construct a slice of maps](https://github.com/isalevine/showme/blob/8409ba6eb1357f3726817c71f6bf7117ec730a60/showme_test.go#L70)

### 3. This is a recreation of the core functionality of my first code-school project, [`Show Randomizer`](https://github.com/isalevine/show-randomizer)!
I had several goals with recreating the TV-show-finding functionality of `showme`:
* Build it in fewer lines of code.
* Make the code as readable as possible.
* Build it with test coverage.
* Condense the original interface into a single console command, i.e. `showme "30 rock"`.
I'm proud to say I accomplished all of those goals in the initial version of the code!


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

Alternately, you can add `alias showme="go run ~/coding/golang/showme/showme.go"` to your bash profile (or `alias showme="go run ~/coding/golang/showme/showme"` for the binary). This will allow you to execute `showme "tv show title"` from anywhere in your console.

**Note that at this time, title names should be exact--including punctuation!** For example, `showme "bob's burgers"` will work, but `showme "bobs burgers` (without the apostrophe) will not.
A future update will print out a list of suggested titles if an exact match is not found.


## Examples

### Running Script
Before building binary:
```
$ go run showme.go "30 rock"

// Output:

OK! From the show '30 Rock', you should watch:

Season 4, Episode 22 - I Do Do

Enjoy!
```

After adding `alias showme="go run ~/coding/golang/showme/showme.go"` to bash profile:
```
$ showme "30 rock"

// Output:

OK! From the show '30 Rock', you should watch:

Season 3, Episode 11 - St. Valentine's Day

Enjoy!
```

### Running Binary
After building binary:
```
$ ./showme "30 rock"

// Output:

OK! From the show '30 Rock', you should watch:

Season 6, Episode 15 - The Shower Principle

Enjoy!
```

After adding `alias showme="go run ~/coding/golang/showme/showme"` to bash profile:
```
$ showme "30 rock"

// Output:

OK! From the show '30 Rock', you should watch:

Season 3, Episode 14 - The Funcooker

Enjoy!
```

# Notes
## Planned Additions:
* When multiple TV show titles are found, print a list of titles to the console to help users re-search for the correct show.
* Add a flag to set the number of episodes returned.
* Add option to pass title argument via `title=tv_show_name` (code already handles parsing underscores in title)
* Add functionality to create aliases for quicker searching, i.e. alias `koth` for `King of the Hill`
