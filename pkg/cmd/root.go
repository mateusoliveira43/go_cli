package cmd

import (
	"fmt"
	"os"

	"github.com/mateusoliveira43/go_cli/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	dataFile string
	cfgFile  string
	debug    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "go_cli",
	Long: "Manage To Do items",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		util.Fatal(fmt.Sprintf("%v", err))
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	pwd, err := os.Getwd()
	if err != nil {
		util.Error("Unable to detect current directory")
		util.Fatal(fmt.Sprintf("%v", err))
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug")
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", pwd+string(os.PathSeparator)+"db.json", "data file to store todos")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

// Read in config file and ENV variables if set.
func initConfig() {
	// TODO
	viper.SetConfigName(".tri")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tri")

	// if a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if configFile := viper.ConfigFileUsed(); len(configFile) > 0 {
			util.Info("Using config file:" + configFile)
		}
		util.Fatal("No environment variable or config file provided for --config flag")
	}
}

func printFlagsValues(f *pflag.Flag) {
	util.Debug(fmt.Sprintf("\t%v: %v", f.Name, f.Value))
}

func DebugFlagsValues(cmd *cobra.Command, args []string) {
	if debug {
		util.Debug("Flags values:")
		cmd.Flags().VisitAll(printFlagsValues)
	}
}
