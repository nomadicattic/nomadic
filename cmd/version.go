package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	versionCmd.Flags().BoolVarP(&Short, "short", "", false, "Short form of version")
	rootCmd.AddCommand(versionCmd)
}

var Version string = "0.0.0"
var Hash string = "n/a"
var Short bool = false

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		if Short {
			fmt.Printf("%s\n", Version)
			return
		}
		fmt.Printf("%s (%s)\n", Version, Hash)
	},
}
