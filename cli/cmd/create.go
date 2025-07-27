/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli/internal/database"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
func createCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples`,
		RunE:  runCreate(categoryDb),
	}
}

func runCreate(categoryDb database.Category) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// Get the flags for name and description
		// If name or description is not provided, it will return help function
		if len(args) < 2 {
			return cmd.Help()
		}
		// If name or description is empty, it will return help function
		if !cmd.Flags().Changed("name") || !cmd.Flags().Changed("description") {
			return cmd.Help()
		}
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDb.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := createCmd(GetCategoryDb(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
