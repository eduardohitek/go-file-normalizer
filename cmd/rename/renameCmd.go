package rename

import (
	"github.com/spf13/cobra"
)

var RenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files",
	Long:  `rename`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return renameFiles(path)
	},
}

func init() {
	RenameCmd.Flags().StringVar(&path, "path", "", "Path")
	RenameCmd.MarkFlagRequired("path")
}
