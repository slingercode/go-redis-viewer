package cmd

import (
	"github.com/slingercode/redis-viewer/pkg"
	"github.com/slingercode/redis-viewer/views"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
  Use: 		"inspect",
  Short: 	"Start the inspect job",
  Long: 	"Start the inspect job",
  Run: func(cmd *cobra.Command, args []string) {
    rclient := pkg.ConnectRedis(URI, PORT)

    views.SetupWindow(rclient)
  },
}

func init() {
  rootCmd.AddCommand(initCmd)
}
