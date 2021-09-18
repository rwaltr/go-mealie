package model

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
