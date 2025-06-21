package main

import (
	"fmt"
	"goapp/cmd/cli/commands"
	"goapp/config"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cfg := newConfig("./config/config.yml")

	root := &cobra.Command{
		Use:   "goapp-cli",
		Short: "Goapp console",
	}

	commands.InitMakeMigration(root, cfg)
	commands.InitMigrate(root, cfg)
	commands.InitMigrateDown(root, cfg)
	commands.InitMigrateStatus(root, cfg)

	err := root.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newConfig(file string) *config.Config {
	instance, err := config.New(file)
	if err != nil {
		panic(err)
	}

	instance.Logger.Level = "Error"
	instance.Logger.Encoding = "pretty"

	return instance
}
