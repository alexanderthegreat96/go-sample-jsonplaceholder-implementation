# {JSON} Placeholder - Go Implementation

A small Go project demonstrating how to use **goroutines** and **channels** to perform **parallel API requests** using the [JSONPlaceholder API](https://jsonplaceholder.typicode.com/).

This example fetches **posts** and their corresponding **comments** concurrently using worker goroutines — making it a great intro to parallel computing in Go.

---

## 🚀 Features

- Concurrent fetching of posts and comments
- Basic worker pool setup with `sync.WaitGroup`
- Channels for communication between goroutines
- Defensive coding with error handling and nil checks
- Clean and readable output formatting

---

## 📡 Used Endpoints

- `GET /posts`
- `GET /posts/{id}/comments`

All endpoints are powered by [JSONPlaceholder](https://jsonplaceholder.typicode.com), a free fake REST API for testing and prototyping.