package main

import (
	"github.com/246859/goplay"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var AppVersion string

var (
	Address     string
	Proxy       string
	Timeout     time.Duration
	ShowVersion bool
)

var rootCmd = &cobra.Command{
	SilenceUsage: true,
	Use:          "goplay",
	Long: "cmd tool to interact with go playground server,\n" +
		"see https://github.com/246859/goplay to learn more about goplay",
	RunE: func(cmd *cobra.Command, args []string) error {
		if ShowVersion {
			if len(AppVersion) == 0 {
				os.Stdout.WriteString("v1.1.1")
			} else {
				os.Stdout.WriteString(AppVersion)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&ShowVersion, "version", "v", false, "show goplay version")

	rootCmd.PersistentFlags().StringVarP(&Address, "address", "d", goplay.DefaultPlayground, "specified the go playground address")
	rootCmd.PersistentFlags().StringVarP(&Proxy, "proxy", "p", "", "proxy address")
	rootCmd.PersistentFlags().DurationVarP(&Timeout, "timeout", "t", time.Second*20, "http request timeout")
	// subcommands
	rootCmd.AddCommand(
		VersionCmd,
		HealthCheckCmd,
		ViewCmd,
		ShareCmd,
		FmtCmd,
		CompileCmd,
	)
}

func NewClient() (*goplay.Client, error) {
	client, err := goplay.NewClient(goplay.Options{
		Address: Address,
		Proxy:   Proxy,
		Timeout: Timeout,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	_ = rootCmd.Execute()
}
