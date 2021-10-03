/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/rwaltr/go-mealie/pkg/client"
	"github.com/rwaltr/go-mealie/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Send a URL to Mealie to scrape and add to your recipebook",
	Long: `Send a URL to Mealie to scrape and add to your recipebook`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.ReadInConfig()
		config, err := utils.LoadConfig()
		if err != nil {
			fmt.Println(err)
		}
		c := client.InitClient(&config)
		response, err := c.Scrapeurl(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(response)
		if preview == true {
			response, err := c.GetRecipe(response)
			if err != nil {
				fmt.Println(err)
			}
			utils.PrettyViewRecipe(response)
		}

	},
}
var preview bool

func init() {
	rootCmd.AddCommand(scrapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scrapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scrapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	scrapeCmd.Flags().BoolVarP(&preview, "preview", "p", false, "Preview scraped recipe")

}
