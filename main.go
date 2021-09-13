package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	mealieURL := os.Getenv("MEALIE_URL")
	if mealieURL == "" {
		fmt.Println("Mealie URL is required")
		os.Exit(1)
	}
	fmt.Println("The URL is", mealieURL)
	mealietoken := os.Getenv("MEALIE_TOKEN")
	if mealietoken == "" {
		fmt.Println("Mealie token is required")
		os.Exit(1)
	}

	response, err := sendreq(mealieURL+"/api/debug", mealietoken, "")
	if err != nil {
		fmt.Println("Request error:", err)
		os.Exit(1)
	}
	var results map[string]interface{}
	json.Unmarshal(response, &results)

	fmt.Println("Production is", results["production"])

	response, err = sendreq(mealieURL+"/api/recipes/summary?start=0&limit=9999", mealietoken, "")
	if err != nil {
		fmt.Println("Request error:", err)
		os.Exit(1)
	}
	json.Unmarshal([]byte(response), &results)
	fmt.Println("results is", results)
}

func sendreq(url string, token string, body string) ([]byte, error) {

	digestedbody := strings.NewReader(body)

	req, err := http.NewRequest("GET", url, digestedbody)
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//var result map[string]interface{}
	//json.Unmarshal([]byte(responseData), &result)

	return responseData, nil
}
