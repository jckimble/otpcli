/*
Copyright © 2019 James Kimble <jckimble@pm.me>

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
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"path/filepath"

	"bufio"
	"bytes"
	"github.com/pquerna/otp"
	"io"
	"os/exec"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "otpcli",
	Short: "A brief description of your application",
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/otpcli/config.yaml)")
	rootCmd.PersistentFlags().StringP("in", "i", "", "file containing otp secrets")
	viper.BindPFlag("otpsecrets", rootCmd.PersistentFlags().Lookup("in"))

	rootCmd.PersistentFlags().StringP("gpg", "g", "", "gpg key to use for decryption")
	viper.BindPFlag("gpgkey", rootCmd.PersistentFlags().Lookup("gpg"))
}

func loadKeys(cmd *cobra.Command) ([]*otp.Key, error) {
	var reader *bufio.Reader
	secrets := viper.GetString("otpsecrets")
	if secrets == "" || secrets == "-" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		var err error
		secrets, err = homedir.Expand(secrets)
		if err != nil {
			return nil, err
		}
		if gpg := viper.GetString("gpgkey"); gpg != "" {
			var out bytes.Buffer
			cmd := exec.Command("gpg", "-q", "--armor", "--decrypt", "-r", gpg, secrets)
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				return nil, err
			}
			reader = bufio.NewReader(&out)
		} else {
			file, err := os.Open(secrets)
			if err != nil {
				return nil, err
			}
			defer file.Close()
			reader = bufio.NewReader(file)
		}
	}
	keyarr := []*otp.Key{}
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		key, err := otp.NewKeyFromURL(string(line))
		if err != nil {
			return nil, err
		}
		keyarr = append(keyarr, key)
	}
	return keyarr, nil
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configDir := filepath.Join(home, ".config/otpcli")
		viper.AddConfigPath(configDir)
		viper.SetConfigName("config")
		bareSecrets := filepath.Join(configDir, "secrets.txt")
		if _, err := os.Stat(bareSecrets); err == nil {
			viper.SetDefault("otpsecrets", bareSecrets)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()
}
