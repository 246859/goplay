package main

import (
	"errors"
	"github.com/spf13/cobra"
	"os"
	"unsafe"
)

var (
	FixImport bool
	WithVet   bool
	SnippetId string
)

var FmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "fmt code snippet",
	Args:  cobra.MaximumNArgs(1),
	RunE:  DoFmt,
}

var CompileCmd = &cobra.Command{
	Use:   "compile",
	Short: "compile and run code snippet",
	Args:  cobra.MaximumNArgs(1),
	RunE:  DoCompile,
}

func init() {
	FmtCmd.Flags().BoolVar(&FixImport, "fix", false, "fix imports while fmt")
	FmtCmd.Flags().StringVarP(&File, "file", "f", "", "specified file to fmt")
	FmtCmd.Flags().StringVarP(&SnippetId, "id", "i", "", "specified id of code snippet")
	CompileCmd.Flags().BoolVar(&WithVet, "vet", false, "vet before compile")
	CompileCmd.Flags().StringVarP(&File, "file", "f", "", "specified file to compile")
	CompileCmd.Flags().StringVarP(&SnippetId, "id", "i", "", "specified id of code snippet")
}

func DoFmt(cmd *cobra.Command, args []string) error {
	if len(args) == 0 && len(File) == 0 && len(SnippetId) == 0 {
		return errors.New("no code snippet provided")
	}

	client, err := NewClient()
	if err != nil {
		return err
	}

	var content []byte
	if len(args) > 0 {
		s := args[0]
		content = unsafe.Slice(unsafe.StringData(s), len(s))
	} else if len(File) > 0 {
		bytes, err := os.ReadFile(File)
		if err != nil {
			return err
		}
		content = bytes
	} else {
		bytes, err := client.View(SnippetId)
		if err != nil {
			return err
		}
		content = bytes
	}

	fmtRaw, err := client.FmtRaw(content, FixImport)
	if err != nil {
		return err
	}
	os.Stdout.Write(fmtRaw)
	return nil
}

func DoCompile(cmd *cobra.Command, args []string) error {
	if len(args) == 0 && len(File) == 0 && len(SnippetId) == 0 {
		return errors.New("no code snippet provided")
	}

	client, err := NewClient()
	if err != nil {
		return err
	}

	var content []byte
	if len(args) > 0 {
		s := args[0]
		content = unsafe.Slice(unsafe.StringData(s), len(s))
	} else if len(File) > 0 {
		bytes, err := os.ReadFile(File)
		if err != nil {
			return err
		}
		content = bytes
	} else {
		bytes, err := client.View(SnippetId)
		if err != nil {
			return err
		}
		content = bytes
	}

	result, err := client.CompileRaw(content, WithVet)
	if err != nil {
		return err
	}
	if result.Errors != "" {
		return errors.New(result.Errors)
	}
	for _, event := range result.Events {
		os.Stdout.WriteString(event.Message)
	}
	return nil
}
