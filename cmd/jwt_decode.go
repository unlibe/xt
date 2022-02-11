
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
	Use:   "jwt:decode",
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
}
