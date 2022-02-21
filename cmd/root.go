package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nomadic",
	Short: "nomad in containers cli",
}

func Execute(v string, h string) error {
	Version = v
	Hash = h
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
