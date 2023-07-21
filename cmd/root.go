package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dataFile string
	cfgFile  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go_cli",
	Short: "TODO where does this appear?",
	Long:  "TODO description",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// home, err := homedir.Dir()
	// if err != nil {
	// 	log.Println("Unable to detect home directory. Please set data file using --datafile")
	// }
	pwd, err := os.Getwd()
	if err != nil {
		log.Println("Unable to detect PWD")
	}
	// log.Println("PWD", pwd)

	// TODO add debug and create logger

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", pwd+string(os.PathSeparator)+"db.json", "data file to store todos")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// TODO what does this do?
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
			fmt.Println("Using config file:", configFile)
		}
		log.Fatalln("No environment variable or config file provided for --config flag")
	}
}
