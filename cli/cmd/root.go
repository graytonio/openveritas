package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

type Config struct {
	Host         string
	ConfigFile   string
	ServerSelect string
	Verbose      bool
	Basic        bool
}

var detailed bool

var (
	config Config
	logger *log.Logger

	rootCmd = &cobra.Command{
		Use:     "veritas",
		Short:   "A CLI for interacting with a veritas server",
		Version: "0.4",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	logger = log.New(ioutil.Discard, "", log.LstdFlags)
	home, _ := os.UserHomeDir()
	cobra.OnInitialize(initConfig)

	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s:" .}}{{end}}{{printf "%s" .Version}}`)

	rootCmd.PersistentFlags().StringP("config", "c", path.Join(home, ".veritasrc"), "config file")
	rootCmd.PersistentFlags().StringVarP(&config.ServerSelect, "server", "s", "default", "server configuration to use")
	rootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", false, "enable verbose logging")
	rootCmd.PersistentFlags().BoolVarP(&config.Basic, "basic", "b", false, "enable basic output")
}

func initConfig() {
	if config.Verbose {
		logger.SetOutput(os.Stdout)
	}

	cfg, err := ini.Load(rootCmd.Flag("config").Value.String())
	if err != nil {
		fmt.Printf("Failed to Read Config File; %s", err.Error())
		os.Exit(1)
	}

	section := cfg.Section(rootCmd.Flag("server").Value.String())
	config.Host = section.Key("host").Value()
}
