package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yang-zzhong/logf"
	"github.com/yang-zzhong/zzblog"
	"os"
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
	var config string
	start := &cobra.Command{
		Use:   "start",
		Short: "start http service",
		Long:  `start http service`,
		Run: func(cmd *cobra.Command, args []string) {
			zzblog.InitConfig(config)
			c := zzblog.GetConfig()
			logf.SetPath(c.LogPath)
			h := zzblog.NewHttp(c.Root, c.DocRoot)
			panic(h.Start(":" + c.Port))
		},
	}
	start.Flags().StringVarP(&config, "config", "c", "config.yml", "config file")
	rootCmd.AddCommand(start)
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
