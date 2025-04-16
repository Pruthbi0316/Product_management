package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("Monday, 02-Jan-2006 15:04:05 MST")
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Current Date & Time</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					text-align: center;
					margin-top: 100px;
				}
				h1 {
					color: #333;
				}
			</style>
		</head>
		<body>
			<h1>Current Date and Time</h1>
			<p>%s</p>
		</body>
		</html>
	`, currentTime)
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
