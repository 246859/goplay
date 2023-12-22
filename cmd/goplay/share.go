package main

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"unsafe"
)

var (
	File string
)

var ShareCmd = &cobra.Command{
	Use:     "share",
	Short:   "share your code to go playground",
	Long:    "share your code to go playground, use the -f to share specific file, \nuse the -r to share specific raw content.",
	Example: "  goplay share -f main.go",
	Args:    cobra.MaximumNArgs(1),
	RunE:    DoShare,
}

func init() {
	ShareCmd.Flags().StringVarP(&File, "file", "f", "", "specified file to share")
}

func DoShare(cmd *cobra.Command, args []string) error {
	if len(args) == 0 && len(File) == 0 {
		return errors.New("no code snippet provided")
	}

	var content []byte
	if len(args) > 0 {
		s := args[0]
		content = unsafe.Slice(unsafe.StringData(s), len(s))
	} else {
		bytes, err := os.ReadFile(File)
		if err != nil {
			return err
		}
		content = bytes
	}

	client, err := NewClient()
	if err != nil {
		return err
	}

	snippetId, err := client.Share(content)
	if err != nil {
		return err
	}
	os.Stdout.WriteString(snippetId)
	return nil
}
