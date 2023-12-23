package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var (
	FixImport bool
	WithVet   bool
	SnippetId string
)

var FmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "fmt code snippet",
	Example: "  goplay fmt example.go example1.go\n" +
		"  cat example.go | goplay fmt\n" +
		"  goplay fmt T9_8fv9CyRh",
	RunE: DoFmt,
}

var CompileCmd = &cobra.Command{
	Use:   "run",
	Short: "compile and run code snippet in playground",
	Example: "  goplay run example.go example1.go\n" +
		"  cat example.go | goplay run\n" +
		"  goplay run T9_8fv9CyRh",
	RunE: DoCompile,
}

func init() {
	FmtCmd.Flags().BoolVar(&FixImport, "fix", false, "fix imports while fmt")
	FmtCmd.Flags().StringVarP(&SnippetId, "id", "i", "", "specified id of code snippet")
	CompileCmd.Flags().BoolVar(&WithVet, "vet", false, "vet before compile")
	CompileCmd.Flags().StringVarP(&SnippetId, "id", "i", "", "specified id of code snippet")
}

func DoFmt(cmd *cobra.Command, args []string) error {
	var content [][]byte

	client, err := NewClient()
	if err != nil {
		return err
	}

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
	} else if len(SnippetId) > 0 {
		bytes, err := client.View(SnippetId)
		if err != nil {
			return err
		}
		content = append(content, bytes)
	}

	if len(content) == 0 {
		return errors.New("no code snippet provided")
	}

	for i, bytes := range content {
		os.Stdout.WriteString(fmt.Sprintf("#%d\n", i+1))
		fmtRaw, err := client.FmtRaw(bytes, FixImport)
		if err != nil {
			return err
		}
		os.Stdout.Write(fmtRaw)
		if i != len(content)-1 {
			os.Stdout.Write([]byte{'\n'})
		}
	}

	return nil
}

func DoCompile(cmd *cobra.Command, args []string) error {
	var content [][]byte

	client, err := NewClient()
	if err != nil {
		return err
	}

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
	} else if len(SnippetId) > 0 {
		bytes, err := client.View(SnippetId)
		if err != nil {
			return err
		}
		content = append(content, bytes)
	}

	if len(content) == 0 {
		return errors.New("no code snippet provided")
	}

	for i, bytes := range content {
		os.Stdout.WriteString(fmt.Sprintf("#%d\n", i+1))
		result, err := client.CompileRaw(bytes, WithVet)
		if err != nil {
			return err
		}
		if len(result.Errors) > 0 {
			os.Stderr.WriteString(result.Errors)
			continue
		}
		for _, event := range result.Events {
			os.Stdout.WriteString(event.Message)
		}
		if i != len(content)-1 {
			os.Stdout.Write([]byte{'\n'})
		}
	}

	return nil
}
