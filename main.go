package main

import (
	"encoding/json"
	"errors"
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
	mealietoken := os.Getenv("MEALIE_TOKEN")
	if mealietoken == "" {
		fmt.Println("Mealie token is required")
		os.Exit(1)
	}

	// response, err := sendreq(mealieURL+"/api/debug", mealietoken, "")
	// if err != nil {
	// 	fmt.Println("Request error:", err)
	// 	os.Exit(1)
	// }
	// var results map[string]interface{}
	// json.Unmarshal(response, &results)

	// fmt.Println("Production is", results["production"])

	// response, err := sendreq(mealieURL+"/api/recipes/summary?start=0&limit=9999", mealietoken, "")
	// if err != nil {
	// 	fmt.Println("Request error:", err)
	// 	os.Exit(1)
	// }
	// var results map[string]interface{}
	// json.Unmarshal([]byte(response), &results)
	// fmt.Println(string(response))

	// Todo: Upload Recipe Function

	// Todo: Recipe Search function

	// Todo: Recipe View Function

	recipeslug, err := scrapeurl(mealieURL, mealietoken, "https://www.allrecipes.com/recipe/232061/baja-style-fish-tacos")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(recipeslug)
}

func scrapeurl(url string, token string, url2scrape string) (string, error) {

	toscrape := map[string]interface{}{"url": url2scrape}

	requestbody, err := json.MarshalIndent(toscrape, "", "   ")
	if err != nil {
		return "", err
	}
	response, statuscode, err := sendreq(url+"/api/recipes/create-url/", token, "POST", string(requestbody))
	if err != nil {
		return "", err
	}

	if statuscode != 201 {
		return "", errors.New("Non 201 Status")
	}

	return string(response), nil
}

func sendreq(url string, token string, reqtype string, body string) ([]byte, int, error) {

	digestedbody := strings.NewReader(body)

	req, err := http.NewRequest(reqtype, url, digestedbody)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, -1, err
	}

	//var result map[string]interface{}
	//json.Unmarshal([]byte(responseData), &result)

	return responseData, response.StatusCode, nil
}
