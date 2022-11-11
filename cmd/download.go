package cmd

import (
	"math/rand"
	"time"
	"touchgrass/client"
	t "touchgrass/torrent"

	"github.com/spf13/cobra"
)

var peerId [20]byte

var TorrentPath string
var Destination string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download -t [path/to/torrent] -d [download/location]",
	Short: `Downloads the files from the supplied torrent file.`,
	Long:  `Downloads the files from the supplied torrent file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		torrent, err := t.ParseTorrent(TorrentPath)
		if err != nil {
			return err
		}

		_, err = client.Download(peerId, torrent)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.PersistentFlags().StringVarP(
		&TorrentPath, "torrent", "t", "", "A path to the torrent file")
	downloadCmd.PersistentFlags().StringVarP(
		&Destination, "dest", "d", "", "A download destination")

	// generate peer id to identify with
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUWQXYZ1234567890"
	var id [20]byte
	for i := 0; i < 20; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}

	peerId = id
}
