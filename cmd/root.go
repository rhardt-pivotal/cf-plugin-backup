package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"code.cloudfoundry.org/cli/plugin"
)

var (
	cfgFile string

	configJSON string
	target     string
	key        string

	backupDir        string
	backupAppBitsDir string
	backupFile       string

	//CliConnection represents the cf cli connection
	CliConnection     plugin.CliConnection
	backupDestination string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cf-plugin-backup",
	Short: "CloudFoundry CLI Backup plug-in",
	Long:  `CloudFoundry CLI plug-in that allows CF users to backup CF resources through CF API`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cf-plugin-backup.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".cf-plugin-backup") // name of config file (without extension)
	viper.AddConfigPath("$HOME")             // adding home directory as first search path
	viper.AutomaticEnv()                     // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	backupDir = "./"
	backupAppBitsDir = "app-bits"
	backupFile = "cf-backup.json"
}
