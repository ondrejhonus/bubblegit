/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/ondrejhonus/bubblegit/utils"

	"github.com/spf13/cobra"
)

// quickCmd represents the quick command
var quickCmd = &cobra.Command{
	Use:   "quick",
	Short: "Quickly add, commit and push changes to a repository",
	Long: `Quickly add, commit and push changes to a repository.

This command allows you to quickly add, commit and push changes to a repository without 
having to go through the entire process of adding, committing and pushing changes separately. 
It is a convenient way to quickly update a repository with your changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		commitMsg := "Added files"
		if len(args) > 0 {
			commitMsg = ""
			for i := 0; i < len(args); i++ {
				commitMsg += args[i]
				if i != len(args)-1 {
					commitMsg += " "
				}
			}
		}
		output := utils.RunCommand("git", "add", ".")
		output += "Added all files\n"
		output += utils.RunCommand("git", "commit", "-m", commitMsg)
		output += "Commited changes with commit message containing: \"" + commitMsg + "\"\n"
		output += utils.RunCommand("git", "push")
		output += "\nQuick commit complete\n"
		fmt.Println(output)
		// fmt.Println("quick called with msg \"" + commitMsg + "\"")
	},
}

func init() {
	rootCmd.AddCommand(quickCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	quickCmd.PersistentFlags().String("string", "Added files", "A help for commit message")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quickCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
