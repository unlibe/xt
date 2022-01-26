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
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"time"
	"xt/rep"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "xt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var engine *gin.Engine
		gin.DisableConsoleColor()
		//gin.SetMode(gin.ReleaseMode)
		//engine = gin.New()
		engine = gin.Default()
		r := engine

		//jsonToGo := rep.GetJsonToGo()



		// 设置线上静态资源路径
		//handler := fsFunc(func(name string) (fs.File, error) {
		//	assetPath := path.Join("./resource/json-to-go", name)
		//	// If we can't find the asset, fs can handle the error
		//	file, err := jsonToGo.Open(assetPath)
		//	if err != nil {
		//		return nil, err
		//	}
		//	// Otherwise, assume this is a legitimate request routed correctly
		//	return file, err
		//})

		pathHandle := func(dirPath string, assert embed.FS) fsFunc {
			return fsFunc(func(name string) (fs.File, error) {
				assetPath := path.Join(dirPath, name)
				// If we can't find the asset, fs can handle the error
				file, err := assert.Open(assetPath)
				if err != nil {
					return nil, err
				}
				// Otherwise, assume this is a legitimate request routed correctly
				return file, err
			})
		}

		// 获取静态资源
		//r.StaticFS("/json-to-go", http.FS(jsonToGo))
		//r.StaticFS("/json-to-go", http.FS(handler))
		r.StaticFS("/json-to-go", http.FS(pathHandle("./resource/json-to-go", rep.GetJsonToGo())))
		r.StaticFS("/json-format", http.FS(pathHandle("./resource/json-format", rep.GetJsonJsonFormat())))
		srv := &http.Server{
			Addr:           ":" + "81",
			Handler:        engine,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		fmt.Println("serve start")

		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		_ = <-quit

		fmt.Println("Shutdown Server ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("Server Shutdown:", err)
		}
		fmt.Println("Server exiting")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Run() {
	cobra.CheckErr(rootCmd.Execute())
}
