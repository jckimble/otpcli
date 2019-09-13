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
	"github.com/spf13/cobra"

	"fmt"
	"github.com/pquerna/otp/totp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Copy totp code to clipboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		keys, err := loadKeys(cmd)
		if err != nil {
			return err
		}
		for _, key := range keys {
			account := key.AccountName()
			issuer := key.Issuer()
			if len(args) != 0 {
				reqissuer := ""
				reqaccount := ""
				if len(args) == 1 {
					if v, _ := cmd.Flags().GetBool("rofi"); v {
						splt := strings.SplitN(args[0], " \t- ", 2)
						reqissuer = splt[0]
						if len(splt) == 2 {
							reqaccount = splt[1]
						}
					} else {
						reqissuer = args[0]
					}
				} else if len(args) == 2 {
					reqissuer = args[0]
					reqaccount = args[1]
				}
				if issuer == reqissuer && (reqaccount == "" || reqaccount == account) {
					code, err := totp.GenerateCode(key.Secret(), time.Now())
					if err != nil {
						return err
					}
					return clipboard.WriteAll(code)
				}
				continue
			}
			fmt.Printf("%s \t- %s\n", issuer, account)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	codeCmd.Flags().BoolP("rofi", "r", false, "Output in rofi Format")
}
