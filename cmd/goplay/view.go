package main

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
)

var ViewCmd = &cobra.Command{
	Use:     "view",
	Short:   "view the specified code snippet",
	Example: "  goplay view T9_8fv9CyRh",
	RunE:    DoView,
	Args:    cobra.ExactArgs(1),
}

func DoView(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("no code snippet id specified")
	}
	id := args[0]

	client, err := NewClient()
	if err != nil {
		return err
	}
	bytes, err := client.View(id)
	if err != nil {
		return err
	}
	os.Stdout.Write(bytes)
	return nil
}
