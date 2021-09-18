package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	model "github.com/rwaltr/go-mealie/pkg/model"
)

type Client struct {
	config     *model.MealieConfig
	httpClient http.Client
}

func InitClient(c *model.MealieConfig) *Client {

	return &Client{
		config:     c,
		httpClient: http.Client{},
	}
}

func (c *Client) GetHTTP(endpoint string, responsebody interface{}) error {
	// TODO, Make more flexable
	fullUrl := fmt.Sprintf("%s/api/%s", c.config.Url, endpoint)
	req, err := http.NewRequest("GET", fullUrl, nil)
	fmt.Println("Sending GET request to: " + fullUrl)

	req.Header.Set("Authorization", "Bearer "+c.config.Token)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		errMsg := fmt.Sprintf("An Error has Occurred Durring Statuscode %d", response.StatusCode)
		return errors.New(errMsg)
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(responsebody)
}

func (c *Client) PostHTTPGetString(endpoint string, body string) (string, error) {
	// TODO, Make more flexable
	fullUrl := fmt.Sprintf("%s/api/%s", c.config.Url, endpoint)
	req, err := http.NewRequest("POST", fullUrl, strings.NewReader(body))
	fmt.Println("Sending POST request to: " + fullUrl)
	fmt.Println("")

	req.Header.Set("Authorization", "Bearer "+c.config.Token)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(req)

	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		errMsg := fmt.Sprintf("An Error has Occurred Durring Statuscode %d", response.StatusCode)
		return "", errors.New(errMsg)
	}

	defer response.Body.Close()
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	rr := strings.Trim(string(r), "\"")

	return string(rr), nil
}

func (c *Client) Scrapeurl(url2scrape string) (string, error) {

	toscrape := map[string]interface{}{"url": url2scrape}

	requestbody, err := json.Marshal(toscrape)
	if err != nil {
		return "", err
	}
	fmt.Printf("Formmated request for URL:%s,\n\n%s", url2scrape, string(requestbody))
	//response, err := sendreq(url+"/api/recipes/create-url/", token, "POST", string(requestbody))
	response, err := c.PostHTTPGetString("recipes/create-url", string(requestbody))
	if err != nil {
		return "", err
	}

	return string(response), nil
}

func (c *Client) GetRecipe(slug string) (model.Recipe, error) {
	var r model.Recipe
	if err := c.GetHTTP("recipes/"+slug, &r); err != nil {
		return r, err
	}

	return r, nil

}

func (c *Client) AllRecipesSummaries() (model.RecipeSummaries, error) {
	var r model.RecipeSummaries
	if err := c.GetHTTP("recipes/summary?start=0&limit=9999", &r); err != nil {
		return r, err
	}

	return r, nil
}
