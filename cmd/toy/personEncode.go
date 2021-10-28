package main

import (
	"fmt"

	pb "github.com/brunocalza/textile-stack/gen/proto/person"
	"github.com/spf13/cobra"
	"github.com/textileio/cli"
	"google.golang.org/protobuf/proto"
)

var personEncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encode receives the person info and encodes using Protobuffer",
	Long:  `encode receives the person info and encodes using Protobuffer`,
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
		fmt.Printf("%x\n", data)

	},
}
