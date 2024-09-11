package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://zipcloud.ibsnet.co.jp/api/search?zipcode=0790177"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(byteArray))
}
