package createmigrationcmd

import (
	"emreddit/migration_manager"

	"github.com/spf13/cobra"
)

// CreateMigrationCm represents the createMigrationCmd command
var CreateMigration = &cobra.Command{
	Use:   "create",
	Short: "creates new migration file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migration_manager.Init(args[0])
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMigrationCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMigrationCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
