package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	config     *mealieConfig
	httpClient http.Client
}

func InitClient(c *mealieConfig) *Client {

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
