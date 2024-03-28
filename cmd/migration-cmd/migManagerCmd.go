/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package migration_cmd

import (
	"emreddit/migration_manager"

	"github.com/spf13/cobra"
)

// migManagerCmdCmd represents the migManagerCmd command
var MigManagerUp = &cobra.Command{
	Use:   "up",
	Short: " running migration up",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migration_manager.RunUp()
	},
}
var MigManagerDown = &cobra.Command{
	Use:   "down",
	Short: "running migration down",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migration_manager.RunDown()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migManagerCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migManagerCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
