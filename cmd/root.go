package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
  URI   string
  PORT  string
)

var rootCmd = &cobra.Command{
  Use:    "rv",
  Short:  "Golang redis viewer",
  Long:   `A CLI golang based inspector for Redis instances`,
}

func init() {
  rootCmd.PersistentFlags().StringVarP(&URI, "uri", "u", "localhost", "Redis instance URI")
  rootCmd.PersistentFlags().StringVarP(&PORT, "port", "p", "6379", "Redis instance PORT")
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}
