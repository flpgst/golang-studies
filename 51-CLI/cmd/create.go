/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("name:", category)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("exists:", exists)
		id, _ := cmd.Flags().GetInt16("id")
		fmt.Println("id:", id)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chama antes do run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chama depois do run")
	},
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	return fmt.Errorf("ocorreu um erro")
	// },
}

var category string

func init() {
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&category, "name", "n", "", "category name")
	createCmd.Flags().BoolP("exists", "e", false, "category exists")
	createCmd.Flags().Int16P("id", "i", 0, "category id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
