package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/textileio/cli"
	logger "github.com/textileio/go-log/v2"
)

var (
	cliName           = "toy"
	envPrefix         = strings.ToUpper(cliName)
	defaultConfigPath = filepath.Join(os.Getenv("HOME"), "."+cliName)
	log               = logger.Logger(cliName)
	v                 = viper.New()
)

var rootCmd = &cobra.Command{
	Use:   cliName,
	Short: "toy is a project for learning about Textile's tech stack",
	Long:  `toy is a project for learning about Textile's tech stack`,
	Args:  cobra.ExactArgs(0),
}

func main() {
	cli.CheckErr(rootCmd.Execute())
}

func init() {
	configPath := os.Getenv("TOY_PATH")
	if configPath == "" {
		configPath = defaultConfigPath
	}

	cobra.OnInitialize(func() {
		v.SetConfigType("json")
		v.SetConfigName("config")
		v.AddConfigPath(os.Getenv(envPrefix + "_PATH"))
		v.AddConfigPath(configPath)
		if err := initConfigFile(configPath); err != nil {
			log.Infof("config file can't be read, creating one")
		}
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("reading config file: %s", err)
		}
	})

	rootCmd.AddCommand(personCmd)
	cli.ConfigureCLI(v, envPrefix, []cli.Flag{
		{Name: "log-debug", DefValue: false, Description: "Enable debug level log"},
		{Name: "log-json", DefValue: false, Description: "Enable structured logging"},
		{Name: "postgres-uri", DefValue: "", Description: "Postgres connection string"},
	}, rootCmd.PersistentFlags())

	personCmd.AddCommand(personEncodeCmd)
	personEncodeCmd.Flags().Int("id", 0, "person's id")
	personEncodeCmd.Flags().String("name", "", "person's name")
	personEncodeCmd.Flags().String("email", "", "person's email")
	personEncodeCmd.MarkFlagRequired("id")
	personEncodeCmd.MarkFlagRequired("name")

	personCmd.AddCommand(personStoreCmd)
	personStoreCmd.Flags().Int("id", 0, "person's id")
	personStoreCmd.Flags().String("name", "", "person's name")
	personStoreCmd.Flags().String("email", "", "person's email")
	personStoreCmd.MarkFlagRequired("id")
	personStoreCmd.MarkFlagRequired("name")

	personCmd.AddCommand(personListCmd)

}

func initConfigFile(configPath string) error {
	path := filepath.Join(configPath, "config")
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		log.Fatalf("create config file path: %s", err)
	}

	if err := v.WriteConfigAs(path); err != nil {
		log.Fatalf("creating config file: %s", err)
	}

	return nil
}
