package main

import (
	"context"
	"fmt"

	storepkg "github.com/brunocalza/textile-stack/cmd/toy/store"
	"github.com/spf13/cobra"
	"github.com/textileio/cli"
)

var personListCmd = &cobra.Command{
	Use:   "list",
	Short: "list lists all persons info in the database",
	Long:  `list lists all persons info in the database`,
	Args:  cobra.ExactArgs(0),
	PersistentPreRun: func(c *cobra.Command, args []string) {
		cli.ExpandEnvVars(v, v.AllSettings())
		err := cli.ConfigureLogging(v, nil)
		cli.CheckErrf("setting log levels: %v", err)
	},
	Run: func(c *cobra.Command, args []string) {
		ctx := context.Background()

		store, err := storepkg.New(v.GetString("postgres-uri"))
		if err != nil {
			log.Fatal(err)
		}

		people, _, err := store.ListPeople(ctx)
		if err != nil {
			log.Fatal(err)
		}
		for _, person := range people {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n", person.ID, person.Name, person.Email.String)
		}

	},
}
