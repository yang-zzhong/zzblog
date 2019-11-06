package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"zzblog"
)

var rootCmd = &cobra.Command{
	Use:   "zzblog",
	Short: "blog",
	Long:  `blog`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "start http service",
		Long:  `start http service`,
		Run: func(cmd *cobra.Command, args []string) {
			h := zzblog.NewHttp("../test/")
			panic(h.Start(":8080"))
		},
	})
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}