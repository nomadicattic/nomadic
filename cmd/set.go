package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets items",
}
