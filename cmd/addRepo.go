/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addRepoCmd represents the addRepo command
var addRepoCmd = &cobra.Command{
	Use:   "add-repo",
	Short: "Add a trusted repository for installing PowerShell modules",
	Long:  `Add a trusted repository for installing PowerShell modules`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addRepo called")
	},
}

func init() {
	rootCmd.AddCommand(addRepoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addRepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addRepoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
