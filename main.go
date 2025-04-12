package main

import (
	"fmt"
	"sync"

	"github.com/alexanderthegreat96/api/core"
)

// a function that calls and API
// we run the function
// grab the api result and store it
// then we output

// the goal is to send 6 requests at once
// be mindful of api rate limitations

func main() {
	var wg sync.WaitGroup                                  // a wait group to sync our stuff
	var postChan chan core.Worker = make(chan core.Worker) // storage buffer for posts
	var workers int = 6                                    // number of consecutive processes

	for i := 1; i <= workers; i++ {
		wg.Add(1)                                // 1 per worker iteration
		go core.GetPostsWorker(i, &wg, postChan) // call the function
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
