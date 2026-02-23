/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/ondrejhonus/bubblegit/utils"

	"github.com/spf13/cobra"
)

var quickCmd = &cobra.Command{
	Use:   "quick <commit_message>",
	Short: "Quickly add, commit and push changes to a repository",
	Long: `Quickly add, commit and push changes to a repository.

This command allows you to quickly add, commit and push changes to a repository without 
having to go through the entire process of adding, committing and pushing changes separately. 
It is a convenient way to quickly update a repository with your changes.

Exampe usage: 
  bubblegit quick Created new commit function`,
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
	quickCmd.PersistentFlags().String("string", "Added files", "A help for commit message")
}
