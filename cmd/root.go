package cmd

import (
	"github.com/eduardohitek/go-file-normalizer/cmd/rename"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file-rename",
	Short: "File rename",
	Long:  "",
}

func init() {

	rootCmd.AddCommand(rename.RenameCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
