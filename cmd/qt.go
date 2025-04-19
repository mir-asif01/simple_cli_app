/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// qtCmd represents the qt command
var qtCmd = &cobra.Command{
	Use:   "qt",
	Short: "Get a random quote with author name.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: getQuote,
}

func init() {
	rootCmd.AddCommand(qtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Quote struct {
	Id    string `json:"id"`
	Quote string `json:"quote"`
	Autor string `json:"author"`
}

func getQuote(cmd *cobra.Command, args []string) {
	url := "https://dummyjson.com/quotes/random"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := Quote{}
		json.Unmarshal(body, &data)
		fmt.Printf("Quote : %v\n", data.Quote)
		fmt.Printf("Author : %v\n", data.Autor)
	}

}
