package apicmd

import (
	"emreddit/api"

	"github.com/spf13/cobra"
)

// apiCmdCmd represents the apiCmd command
var RunApi = &cobra.Command{
	Use:   "api",
	Short: "runs backend",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api.ListenPort()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
