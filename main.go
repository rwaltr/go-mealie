package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
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

func (c *Client) getRecipe(slug string) (Recipe, error) {
	var r Recipe
	if err := c.GetHTTP("recipes/"+slug, &r); err != nil {
		return r, err
	}

	return r, nil

}

func (c *Client) allRecipesSummaries() (RecipeSummaries, error) {
	var r RecipeSummaries
	if err := c.GetHTTP("recipes/summary?start=0&limit=9999", &r); err != nil {
		return r, err
	}

	return r, nil
}

func grabRecipeDownloadEndpoint(recipe Recipe) string {
	r := fmt.Sprintf("%s/api/recipes/%s/zip")
	return r
}

func prettyViewRecipe(recipe Recipe) error {
	fmt.Printf("# %s\n\n## Description: %s\n\nServings: %s\n\nMealie URL: %s\n\nOriginal Url: %s\n\n",
		recipe.Name,
		recipe.Description,
		recipe.RecipeYield,
		viper.GetString("url")+"/recipe/"+recipe.Slug,
		recipe.OrgURL)
	fmt.Println("## Recipe List")
	for i := range recipe.RecipeIngredient {
		if recipe.RecipeIngredient[i].DisableAmount == false {
			fmt.Printf("%s %s %s %s\n",
				string(recipe.RecipeIngredient[i].Quantity),
				recipe.RecipeIngredient[i].Unit,
				recipe.RecipeIngredient[i].Title,
				recipe.RecipeIngredient[i].Note)
		} else {
			fmt.Printf("- %s\n", string(recipe.RecipeIngredient[i].Note))
		}
	}
	fmt.Printf("\n\n## Instructions\n")
	for i := range recipe.RecipeInstructions {
		fmt.Printf("- %s\n\n", recipe.RecipeInstructions[i].Text)

	}

	return nil
}

func (c *Client) scrapeurl(url2scrape string) (string, error) {

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

type mealieConfig struct {
	Url   string `mapstructure:"url"`
	Token string `mapstructure:"token"`
}

type Recipe struct {
	ID               int           `json:"id"`
	Name             string        `json:"name"`
	Slug             string        `json:"slug"`
	Image            string        `json:"image"`
	Description      string        `json:"description"`
	RecipeCategory   []interface{} `json:"recipeCategory"`
	Tags             []interface{} `json:"tags"`
	Rating           interface{}   `json:"rating"`
	DateAdded        string        `json:"dateAdded"`
	DateUpdated      string        `json:"dateUpdated"`
	RecipeYield      string        `json:"recipeYield"`
	RecipeIngredient []struct {
		Title         interface{} `json:"title"`
		Note          string      `json:"note"`
		Unit          interface{} `json:"unit"`
		Food          interface{} `json:"food"`
		DisableAmount bool        `json:"disableAmount"`
		Quantity      int         `json:"quantity"`
	} `json:"recipeIngredient"`
	RecipeInstructions []struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"recipeInstructions"`
	Nutrition struct {
		Calories            interface{} `json:"calories"`
		FatContent          interface{} `json:"fatContent"`
		ProteinContent      interface{} `json:"proteinContent"`
		CarbohydrateContent interface{} `json:"carbohydrateContent"`
		FiberContent        interface{} `json:"fiberContent"`
		SodiumContent       interface{} `json:"sodiumContent"`
		SugarContent        interface{} `json:"sugarContent"`
	} `json:"nutrition"`
	Tools       []interface{} `json:"tools"`
	TotalTime   string        `json:"totalTime"`
	PrepTime    string        `json:"prepTime"`
	PerformTime string        `json:"performTime"`
	Settings    struct {
		Public          bool `json:"public"`
		ShowNutrition   bool `json:"showNutrition"`
		ShowAssets      bool `json:"showAssets"`
		LandscapeView   bool `json:"landscapeView"`
		DisableComments bool `json:"disableComments"`
		DisableAmount   bool `json:"disableAmount"`
	} `json:"settings"`
	Assets []interface{} `json:"assets"`
	Notes  []interface{} `json:"notes"`
	OrgURL string        `json:"orgURL"`
	Extras struct {
	} `json:"extras"`
	Comments []interface{} `json:"comments"`
}

// TODO Use config object
func main() {

	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
	}

	c := InitClient(&config)
	if err != nil {
		fmt.Println(err)
	}

	resultSlug, err := c.scrapeurl("https://www.tasteofhome.com/recipes/frito-pie/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultSlug)

	resultRecipe, err := c.getRecipe(resultSlug)
	if err != nil {
		fmt.Println(err)
	}
	prettyViewRecipe(resultRecipe)

	// var testRecipe Recipe
	// if err := c.GetHTTP("recipes/keto-chicken-nuggets", &testRecipe); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(testRecipe.Name)

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

	// myrecipe, err := grabRecipe(mealieURL, mealietoken, "loaded-smashed-potatoes")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// prettyViewRecipe(myrecipe)

}

func loadConfig() (mealieConfig, error) {
	var result mealieConfig
	viper.AddConfigPath("$XDG_CONFIG_HOME/go-mealie")
	viper.AddConfigPath("$HOME/.config/go-mealie")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("MEALIE")
	viper.BindEnv("url")
	viper.BindEnv("token")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return result, err
	}

	err = viper.Unmarshal(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
