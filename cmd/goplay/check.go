package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "get go version of the playground server",
	RunE:  DoVersion,
}

func DoVersion(cmd *cobra.Command, args []string) error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	version, err := client.Version()
	if err != nil {
		return err
	}
	os.Stdout.WriteString(fmt.Sprintf("name: %s\nrelease: %s\nversion: %s\n", version.Name, version.Release, version.Version))
	return nil
}

var HealthCheckCmd = &cobra.Command{
	Use:   "health",
	Short: "check whether the playground server is healthy",
	RunE:  DoHealthCheck,
}

func DoHealthCheck(cmd *cobra.Command, args []string) error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	healthy, err := client.HealCheck()
	if err != nil {
		return err
	}
	if !healthy {
		return errors.New("health check failed")
	}
	os.Stdout.WriteString("ok")
	return nil
}
