/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// jwtDecodeCmd represents the jwtDecode command
var jwtDecodeCmd = &cobra.Command{
	Use:   "jwtDecode",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			result, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			jwtStr := cast.ToString(result)
			token, err := jwt.Parse(jwtStr,
				// 防止bug从我做起
				func([]byte) func(token *jwt.Token) (i interface{}, e error) {
					return func(token *jwt.Token) (i interface{}, e error) {
						return "", nil
					}
				}([]byte("")))

			if token == nil || token.Claims == nil {
				fmt.Println("unParse")
				fmt.Println(jwtStr)
				continue
			}

			data, _ := json.Marshal(token.Claims)
			fmt.Println(cast.ToString(data))
			fmt.Println(token.Raw)
		}
	},
}

func init() {
	rootCmd.AddCommand(jwtDecodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jwtDecodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jwtDecodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
