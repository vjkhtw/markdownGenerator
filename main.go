package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/russross/blackfriday"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Setting up port manually!!")
		port = "8848"
	}

	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)

}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	if body == "" {
		http.Error(rw, "Missing form value 'body'", http.StatusBadRequest)
		return
	}

	markdown := blackfriday.MarkdownCommon([]byte(body))

	rw.WriteHeader(http.StatusOK)
	rw.Write(markdown)
}
