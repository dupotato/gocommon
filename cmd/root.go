// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	logger "github.com/dupotato/gocommon/dlogger/zlog"
)

var cfgFile string
var printVersion bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  `详细介绍`,
	Run: func(cmd *cobra.Command, args []string) {
		rootserver()
	},
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
	cobra.OnInitialize(loadConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/goapg.yaml)")
}

func loadConfig() {
	initConfig()
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find current dir
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal().Err(err).Msgf("get current path failed")
			os.Exit(1)
		}

		viper.AddConfigPath(dir)
		viper.SetConfigName("aihub-proxy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {

		logger.Infof("Using config file: %v", viper.ConfigFileUsed())
	} else {
		logger.Errorf("Using config file: %v", err)
	}
}
