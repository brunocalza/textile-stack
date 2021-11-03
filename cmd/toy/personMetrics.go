package main

import (
	"context"

	storepkg "github.com/brunocalza/textile-stack/cmd/toy/store"
	"github.com/brunocalza/textile-stack/common"
	"github.com/spf13/cobra"
	"github.com/textileio/cli"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var personMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "metrics is a daemon that collects metrics from database and export to Prometheus",
	Long:  `metrics is a daemon that collects metrics from database and export to Prometheus`,
	Args:  cobra.ExactArgs(0),
	PersistentPreRun: func(c *cobra.Command, args []string) {
		cli.ExpandEnvVars(v, v.AllSettings())
		err := cli.ConfigureLogging(v, nil)
		cli.CheckErrf("setting log levels: %v", err)
	},
	Run: func(c *cobra.Command, args []string) {
		err := common.SetupInstrumentation(v.GetString("metrics-port"))
		if err != nil {
			cli.CheckErrf("failed to setup instrumentation: %v", err)
		}

		store, err := storepkg.New(v.GetString("postgres-uri"))
		if err != nil {
			log.Fatal(err)
		}

		var Meter = metric.Must(global.Meter("toy"))
		Meter.NewInt64GaugeObserver("toy.person.count", func(ctx context.Context, r metric.Int64ObserverResult) {
			people, _, err := store.ListPeople(ctx)
			if err != nil {
				log.Fatal(err)
			}

			r.Observe(int64(len(people)))
		})

		cli.HandleInterrupt(func() {

		})

	},
}
