package main

import "github.com/spf13/cobra"

var personCmd = &cobra.Command{
	Use:   "person",
	Short: "person is a command for executions subcommands in the context of a person",
	Long:  `person is a command for executions subcommands in the context of a person`,
	Args:  cobra.ExactArgs(0),
}
