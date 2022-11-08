package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "touchgrass",
	Short: "A CLI BitTorrent client",
	Long: `A CLI BitTorrent client
TODO:
 - Multi file support
 - Seeding
 - DHT and UDP peer discovery
 - Use the app as a daemon`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
