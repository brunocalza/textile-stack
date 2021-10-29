package main

import (
	"context"

	storepkg "github.com/brunocalza/textile-stack/cmd/toy/store"
	pb "github.com/brunocalza/textile-stack/gen/proto/person"
	"github.com/spf13/cobra"
	"github.com/textileio/cli"
	"google.golang.org/protobuf/proto"
)

var personStoreCmd = &cobra.Command{
	Use:   "store",
	Short: "store stores the person info in the database",
	Long:  `store stores the person info in the database`,
	Args:  cobra.ExactArgs(0),
	PersistentPreRun: func(c *cobra.Command, args []string) {
		cli.ExpandEnvVars(v, v.AllSettings())
		err := cli.ConfigureLogging(v, nil)
		cli.CheckErrf("setting log levels: %v", err)
	},
	Run: func(c *cobra.Command, args []string) {
		id := v.GetInt32("id")
		name := v.GetString("name")
		email := v.GetString("email")

		person := pb.Person{Name: name, Id: id}
		if email != "" {
			person.Email = &email
		}

		data, err := proto.Marshal(&person)
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		store, err := storepkg.New(v.GetString("postgres-uri"))
		if err != nil {
			log.Fatal(err)
		}

		err = store.CreatePerson(ctx, person, data)
		if err != nil {
			log.Fatal(err)
		}

	},
}
