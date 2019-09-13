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

	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/spf13/cobra"

	"image/png"
	"os"
)

// qrsecretCmd represents the qrsecret command
var qrsecretCmd = &cobra.Command{
	Use:   "qrsecret",
	Short: "Outputs a qrcode for adding to other devices",
	RunE: func(cmd *cobra.Command, args []string) error {
		keys, err := loadKeys(cmd)
		if err != nil {
			return err
		}
		for _, key := range keys {
			account := key.AccountName()
			issuer := key.Issuer()
			if len(args) != 0 {
				if issuer == args[0] && (len(args) == 1 || args[1] == account) {
					if out, _ := cmd.Flags().GetString("out"); out != "" && out != "-" {
						outfile, err := os.Create(out)
						if err != nil {
							return err
						}
						defer outfile.Close()
						img, err := key.Image(200, 200)
						if err != nil {
							return err
						}
						if err := png.Encode(outfile, img); err != nil {
							return err
						}
						return nil
					} else {
						qrcodeTerminal.New().Get(key.String()).Print()
						return nil
					}
				}
				continue
			}
			fmt.Printf("%s \t- %s\n", issuer, account)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(qrsecretCmd)
	qrsecretCmd.Flags().StringP("out", "o", "", "Out file for image")
}
