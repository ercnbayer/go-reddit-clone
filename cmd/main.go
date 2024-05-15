package main

import (
	createmigrationcmd "emreddit/cmd/create-migration-cmd"
	migration_cmd "emreddit/cmd/migration-cmd"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "emreddit",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(migration_cmd.MigManagerUp)
	rootCmd.AddCommand(migration_cmd.MigManagerDown)
	rootCmd.AddCommand(createmigrationcmd.CreateMigration)
}
func main() {

	addSubCommands()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Execute()
}
