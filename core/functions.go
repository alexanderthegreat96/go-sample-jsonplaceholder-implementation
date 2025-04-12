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

func getPostsWorker(id int, wg *sync.WaitGroup, postsChan chan<- Worker) {
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

func PostsGoroutine() {
	var wg sync.WaitGroup                        // a wait group to sync our stuff
	var postChan chan Worker = make(chan Worker) // storage buffer for posts
	var workers int = 6                          // number of consecutive processes

	for i := 1; i <= workers; i++ {
		wg.Add(1)                           // 1 per worker iteration
		go getPostsWorker(i, &wg, postChan) // call the function
		// time.Sleep(1 * time.Second) // introduce delays if needed
	}

	// we wait for all goroutines to finish
	// and close the channel which stores our posts
	go func() {
		wg.Wait()
		close(postChan)
	}()

	for workerData := range postChan {
		if len(workerData.Posts) == 0 {
			fmt.Printf("⚠️  Worker %d returned no posts.\n", workerData.WorkerId)
			continue
		}

		fmt.Printf("========== Worker %d ==========\n", workerData.WorkerId)

		for _, post := range workerData.Posts {
			fmt.Printf("\n📌 Post ID: %d\n", post.Id)
			fmt.Printf("👤 User ID: %d\n", post.UserId)
			fmt.Printf("📝 Title  : %s\n", post.Title)

			if len(post.Comments) == 0 {
				fmt.Println("💬 No comments for this post.")
			} else {
				fmt.Printf("💬 Comments (%d):\n", len(post.Comments))
				for _, comment := range post.Comments {
					fmt.Printf("  └─ 💬 Comment ID: %d\n", comment.ID)
					fmt.Printf("     🧑 Name : %s\n", comment.Name)
					fmt.Printf("     📧 Email: %s\n", comment.Email)
					fmt.Printf("     ✏️  Text : %s\n\n", comment.Body)
				}
			}
		}

		fmt.Println("=================================\n")
	}

}
