package utils

import (
	"fmt"

	model "github.com/rwaltr/go-mealie/pkg/model"
	"github.com/spf13/viper"
)

func GrabRecipeDownloadEndpoint(recipe model.Recipe) string {
	r := fmt.Sprintf("%s/api/recipes/%s/zip")
	return r
}

func PrettyViewRecipe(recipe model.Recipe) error {
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
