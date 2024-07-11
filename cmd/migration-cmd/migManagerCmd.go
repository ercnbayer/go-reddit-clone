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

var RunSingleMig = &cobra.Command{
	Use:   "single",
	Short: "running migration up",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migration_manager.RunUpMigration(args[0])
	},
}

func init() {

}
