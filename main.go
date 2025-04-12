package main

import "github.com/alexanderthegreat96/api/core"

// a function that calls and API
// we run the function
// grab the api result and store it
// then we output

// the goal is to send 6 requests at once
// be mindful of api rate limitations

func main() {
	// grabbing posts
	core.PostsGoroutine()
}
