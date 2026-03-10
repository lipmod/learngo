package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(len(body))
	return nil
}

func main() {
	if err := fetch("https://www.lipcoder.top"); err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	}
}
