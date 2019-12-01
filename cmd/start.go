package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	var docroot, root, port string
	start := &cobra.Command{
		Use:   "start",
		Short: "start http service",
		Long:  `start http service`,
		Run: func(cmd *cobra.Command, args []string) {
			h := zzblog.NewHttp(root, docroot)
			panic(h.Start(":" + port))
		},
	}
	start.Flags().StringVarP(&docroot, "docroot", "d", "", "html root")
	start.Flags().StringVarP(&root, "root", "r", "/Users/yangzhong/dev/go/src/zzblog/test/root", "blog root")
	start.Flags().StringVarP(&port, "port", "p", "80", "listen address")
	rootCmd.AddCommand(start)
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
