/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/flpgst/golang-studies/51-CLI/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create new category",
		Long:  "Create new category",
		RunE:  runCreate(categoryDb),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDB.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "category name")
	createCmd.Flags().StringP("description", "d", "", "category description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
