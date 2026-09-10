/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/Vappelquist/quiz/api"
	"github.com/Vappelquist/quiz/store"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the quiz",
	Run: func(cmd *cobra.Command, args []string) {
		s := store.New()
		h := api.NewHandler(s)
		router := api.NewRouter(h)

		addr := ":8080"
		fmt.Println("Quiz API listening on %s\n", addr)
		if err := http.ListenAndServe(addr, router); err != nil {
			fmt.Println("server error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
