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

// listAllRecipiesCmd represents the listAllRecipies command
var listAllRecipiesCmd = &cobra.Command{
	Use:   "listAllRecipies",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.ReadInConfig()
		config, err := utils.LoadConfig()
		if err != nil {
			fmt.Println(err)
		}
		c := client.InitClient(&config)
		response, err := c.AllRecipesSummaries()
		if err != nil {
			fmt.Println(err)
		}
		for i := range response {
			fmt.Println(response[i].Slug)
		}

	},
}

func init() {
	rootCmd.AddCommand(listAllRecipiesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAllRecipiesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAllRecipiesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
