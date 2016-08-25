package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/orus-io/yago/generate"
)

var logger = log.New(os.Stdout, "yago", 0)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "yago",
	Short: "Yet Another GORM",
	Long:  `yago code generator.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.ParseFlags(args)
		output := cmd.Flag("output").Value.String()

		wd, err := os.Getwd()
		if err != nil {
			logger.Fatal(err)
		}
		gofilename := os.Getenv("GOFILE")
		gopackage := os.Getenv("GOPACKAGE")

		if flag := cmd.Flag("package"); flag.Changed {
			gopackage = flag.Value.String()
		}

		err = generate.ProcessFile(logger, wd, gofilename, gopackage, output)
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
	RootCmd.Flags().StringP("output", "o", "", "Set the output file name")
	RootCmd.Flags().StringP("package", "p", "", "Force the package")
}
