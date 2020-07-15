/*
Copyright © 2020 Grant Buskey gbuskey@ncsu.edu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/gbuskey/Jupiter/cmd/forecast"
	"github.com/gbuskey/Jupiter/cmd/history"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Jupiter",
	Short: "Jupiter is a CLI used for weather data",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Run any functions desired on startup
	cobra.OnInitialize(initConfig)

	//Add custom commands
	forecastCmd := forecast.NewForecastCmd()
	rootCmd.AddCommand(forecastCmd)

	historyCmd := history.NewHistoryCmd()
	rootCmd.AddCommand(historyCmd)
}

// initConfig reads in environment variables
func initConfig() {

	// Search config in home directory with name ".Jupiter" (without extension).
	viper.AddConfigPath("./")
	viper.SetConfigName(".Jupiter")

	// Link environment variables to flags
	viper.SetEnvPrefix("JUPITER")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
