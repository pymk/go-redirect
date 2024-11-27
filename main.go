package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type URLStore map[string]string

const (
	portNumber string = ":8080"
	dbPath     string = "data/db.txt"
)

func main() {
	urlStore := make(URLStore)
	readLoadFile(dbPath, urlStore)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirectHandler(w, r, urlStore)
	})

	fmt.Println("localhost" + portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}

func readLoadFile(path string, store URLStore) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		store[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request, store URLStore) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if longURL, ok := store[path]; ok {
		if !strings.HasPrefix(longURL, "http") || !strings.HasPrefix(longURL, "http") {
			longURL = "https://" + longURL
		}
		http.Redirect(w, r, longURL, http.StatusMovedPermanently)
		return
	}
	http.NotFound(w, r)
}
