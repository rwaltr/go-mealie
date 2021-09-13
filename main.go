package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type RecipeSummaries []struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Slug           string   `json:"slug"`
	Image          string   `json:"image"`
	Description    string   `json:"description"`
	RecipeCategory []string `json:"recipeCategory"`
	Tags           []string `json:"tags"`
	Rating         int      `json:"rating"`
	DateAdded      string   `json:"dateAdded"`
	DateUpdated    string   `json:"dateUpdated"`
}

type Recipe struct {
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Image              string   `json:"image"`
	RecipeYield        string   `json:"recipe_yield"`
	RecipeIngredient   []string `json:"recipe_ingredient"`
	RecipeInstructions []struct {
		Text string `json:"text"`
	} `json:"recipe_instructions"`
	Slug           string   `json:"slug"`
	Tags           []string `json:"tags"`
	RecipeCategory []string `json:"recipe_category"`
	Notes          []struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"notes"`
	OrgURL string `json:"org_url"`
	Rating int    `json:"rating"`
	Extras struct {
		Message string `json:"message"`
	} `json:"extras"`
}

// TODO Use config object
func main() {
	viper.SetEnvPrefix("MEALIE")
	viper.BindEnv("url")
	viper.BindEnv("token")
	mealieURL := viper.GetString("url")
	mealietoken := viper.GetString("token")

	if mealieURL == "" {
		fmt.Println("Mealie URL is required")
		os.Exit(1)
	}
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

	// recipeslug, err := scrapeurl(mealieURL, mealietoken, "https://www.allrecipes.com/recipe/17205/eggs-benedict/")
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(recipeslug)

	// results, err := allrecipessummary(mealieURL, mealietoken)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, name := range results {
	// 	fmt.Printf("name: %s\nDescription: %s\nSlug:%s\n\n", name.Name, name.Description, name.Slug)

	// }
	_, err := grabRecipe(mealieURL, mealietoken, "authentic-keto-cornbread")
	if err != nil {
		fmt.Println(err)
	}

}

func csvViewRecipe(recipe Recipe) error {
	return nil
}

func prettyViewRecipe(recipe Recipe) error {
	fmt.Println("Printed Recipe")
	return nil
}

func grabRecipeDownloadendpoint(recipe Recipe) string {
	endpoint := "/api/recipes/" + recipe.Slug + "/zip"
	return endpoint
}

func grabRecipe(url string, token string, recipeslug string) (Recipe, error) {
	var result Recipe
	response, err := sendreq(url+"/api/recipes/"+recipeslug, token, "GET", "")
	if err != nil {
		return result, err
	}

	json.Unmarshal(response, &result)

	return result, nil
}

func allrecipessummary(url string, token string) (RecipeSummaries, error) {

	response, err := sendreq(url+"/api/recipes/summary?start=0&limit=9999", token, "GET", "")
	if err != nil {
		return nil, err
	}
	var results RecipeSummaries
	json.Unmarshal(response, &results)

	return results, nil
}

func scrapeurl(url string, token string, url2scrape string) (string, error) {

	toscrape := map[string]interface{}{"url": url2scrape}

	requestbody, err := json.MarshalIndent(toscrape, "", "   ")
	if err != nil {
		return "", err
	}
	response, err := sendreq(url+"/api/recipes/create-url/", token, "POST", string(requestbody))
	if err != nil {
		return "", err
	}

	return string(response), nil
}

func sendreq(url string, token string, reqtype string, body string) ([]byte, error) {

	digestedbody := strings.NewReader(body)

	req, err := http.NewRequest(reqtype, url, digestedbody)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		errMsg := fmt.Sprintf("An Error has Occurred Durring Statuscode %d", response.StatusCode)
		return nil, errors.New(errMsg)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//var result map[string]interface{}
	//json.Unmarshal([]byte(responseData), &result)

	return responseData, nil
}
