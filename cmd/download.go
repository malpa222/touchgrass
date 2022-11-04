/*
Copyright © 2022 Daniel Lewandowski malpa222@tutanota.com

*/

package cmd

import (
	"log"
	"touchgrass/client/tracker"
	t "touchgrass/torrent"

	"github.com/spf13/cobra"
)

var TorrentPath string
var Destination string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: `Downloads the files from the supplied torrent file.`,
	Long:  `Downloads the files from the supplied torrent file.`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		torrent, err := t.ParseTorrent(TorrentPath)
		if err != nil {
			return err
		}

		peers, err := tracker.GetPeers(torrent)
		log.Println(peers)

		return
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.PersistentFlags().StringVarP(
		&TorrentPath, "torrent", "t", "", "A path to the torrent file")
	downloadCmd.PersistentFlags().StringVarP(
		&Destination, "dest", "d", "", "A download destination")
}
