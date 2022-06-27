/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ssrfCmd represents the ssrf command
var ssrfCmd = &cobra.Command{
	Use:   "ssrf",
	Short: "Simulate SSRF attack.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ssrf called")
	},
}

func init() {
	rootCmd.AddCommand(ssrfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ssrfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ssrfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
