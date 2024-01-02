package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gougou/fetch"
	"gougou/version"
	"os"
)

// Gou is the main command for the dex command-line tool.
var Gou = &cobra.Command{
	Use:           "gougou",
	Long:          "A Cli Tools For Download Some Resource.",
	SilenceErrors: true,
}

// fetchCmd is the command for fetch data.
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch All of NVD Data From Index 0",
	Run: func(cmd *cobra.Command, args []string) {
		fetch.StartFetch()
	},
}

// versionCmd is the command for version.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gougou",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gougou version is " + version.Getversion())
	},
}

func init() {
	Gou.AddCommand(fetchCmd)
	Gou.AddCommand(versionCmd)
}

// run executes the dex command-line tool.
func run() {
	if err := Gou.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// main is the entry point for the dex command-line tool.
func main() {
	run()
}
