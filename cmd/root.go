package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "serato-tools",
	Short: "Serato Tools is a tool to help you perform tasks on Serato",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&musicDir, "dir", "d", "", "Root directory for your music (required)")
	rootCmd.MarkFlagRequired("dir")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}