package main

import (
	"fmt"

	client "github.com/rwaltr/go-mealie/pkg/client"
	util "github.com/rwaltr/go-mealie/pkg/util"
	"github.com/spf13/viper"
)

func main() {

	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
	}

	c := client.InitClient(&config)
	if err != nil {
		fmt.Println(err)
	}

	resultSlug, err := c.Scrapeurl("https://www.allrecipes.com/recipe/17205/eggs-benedict/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultSlug)

	resultRecipe, err := c.GetRecipe(resultSlug)
	if err != nil {
		fmt.Println(err)
	}
	util.PrettyViewRecipe(resultRecipe)

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
