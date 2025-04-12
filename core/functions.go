package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// function that calls the api
func getPosts() ([]Post, error) {
	req, err := http.Get(POSTS_API_URL)

	if err != nil {
		return []Post{}, err
	}

	defer req.Body.Close()

	var posts []Post

	if err := json.NewDecoder(req.Body).Decode(&posts); err != nil {
		return []Post{}, err
	}

	return posts, nil
}

func getPostComments(postId int) ([]Comment, error) {
	req, err := http.Get(fmt.Sprintf("%s/%d/comments", POSTS_API_URL, postId))

	if err != nil {
		return []Comment{}, err
	}

	defer req.Body.Close()

	var comments []Comment

	if err := json.NewDecoder(req.Body).Decode(&comments); err != nil {
		return []Comment{}, nil
	}

	return comments, nil
}

// function that simply sends the request, grabs the result
// and places it inside the channel

func GetPostsWorker(id int, wg *sync.WaitGroup, postsChan chan<- Worker) {
	defer wg.Done() // runs at the end of the code

	// just getting the posts
	posts, err := getPosts()
	if err != nil {
		fmt.Printf("Issue with the posts api request: %s", err.Error())
		return
	}

	// casually grabing the posts for each comment
	for i := range posts {
		comments, err := getPostComments(posts[id].Id)

		if err != nil {
			continue
		}
		posts[i].Comments = comments
	}

	postsChan <- Worker{WorkerId: id, Posts: posts}
}
