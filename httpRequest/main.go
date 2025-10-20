package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("ACCEPT", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(data))

}