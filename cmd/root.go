package cmd

import (
	apicmd "emreddit/cmd/api-cmd"
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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
	rootCmd.AddCommand(apicmd.RunApi)
}
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.emreddit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	addSubCommands()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
