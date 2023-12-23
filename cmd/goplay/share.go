package main

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var ShareCmd = &cobra.Command{
	Use:     "share",
	Short:   "share your code to go playground",
	Long:    "share your code to go playground, use the -f to share specific file, \nuse the -r to share specific raw content.",
	Example: "  goplay share xxx.go xxx.go",
	RunE:    DoShare,
}

func DoShare(cmd *cobra.Command, args []string) error {

	var content [][]byte
	stat, _ := os.Stdin.Stat()
	// pipeline mode
	if stat.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
		all, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		content = append(content, all)
	} else if len(args) > 0 {
		for _, arg := range args {
			bytes, err := os.ReadFile(arg)
			if err != nil {
				return err
			}
			content = append(content, bytes)
		}
	}

	if len(content) == 0 {
		return errors.New("no code snippet provided")
	}

	client, err := NewClient()
	if err != nil {
		return err
	}

	for i, bytes := range content {
		snippetId, err := client.Share(bytes)
		if err != nil {
			return err
		}
		os.Stdout.WriteString(snippetId)
		if i != len(content)-1 {
			os.Stdout.WriteString("\n")
		}
	}
	return nil
}
