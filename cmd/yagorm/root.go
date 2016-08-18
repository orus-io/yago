package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"bitbucket.org/cdevienne/yagorm/generate"
)

var logger = log.New(os.Stdout, "yagorm", 0)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "yagorm",
	Short: "Yet Another GORM",
	Long:  `yagorm code generator.`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			logger.Fatal(err)
		}
		gofilename := os.Getenv("GOFILE")
		gopackage := os.Getenv("GOPACKAGE")

		err = generate.ProcessFile(logger, wd, gofilename, gopackage)
		if err != nil {
			logger.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().BoolP("debug", "d", false, "Enable debug logging")
}
