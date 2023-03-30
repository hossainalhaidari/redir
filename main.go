package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const RedirectsFile = "_redirects"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		value, err := getKey(r.URL.Path[1:])

		if err == nil && value != "" {
			http.Redirect(w, r, value, http.StatusFound)
		}

		fmt.Fprintf(w, "Not Found")
	})
	http.ListenAndServe(":3000", nil)
}

func getKey(key string) (string, error) {
	key = "/" + key + " "

	file, err := os.Open(RedirectsFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, key) {
			return text[len(key):], scanner.Err()
		}
	}

	return "", scanner.Err()
}
