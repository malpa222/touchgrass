/*
Copyright © 2022 Daniel Lewandowski lewandowski-daniel@protonmail.com

*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: `Downloads the files from the supplied torrent file.`,
	Long:  `Downloads the files from the supplied torrent file.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("download called")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	downloadCmd.PersistentFlags().StringP("torrent", "t", "", "A path to the torrent file")
	downloadCmd.PersistentFlags().StringP("dest", "d", "", "A download destination")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
