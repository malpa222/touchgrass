package cmd

import (
	"log"
	"math/rand"
	"time"
	"touchgrass/client"
	"touchgrass/client/tracker"
	t "touchgrass/torrent"

	"github.com/spf13/cobra"
)

var peerId [20]byte

var TorrentPath string
var Destination string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download [path/to/torrent] [download/location]",
	Short: `Downloads the files from the supplied torrent file.`,
	Long:  `Downloads the files from the supplied torrent file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		torrent, err := t.ParseTorrent(TorrentPath)
		if err != nil {
			return err
		}

		peers, err := tracker.GetPeers(peerId, torrent)
		if err != nil {
			return err
		}

		connectToPeers(*peers, torrent.InfoHash)

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

func connectToPeers(peers []tracker.Peer, infoHash [20]byte) {
	var clients []*client.Client
	for _, p := range peers {
		c := client.New(p, infoHash, peerId)

		if err := c.Connect(&p); err != nil {
			log.Println(err)
			continue
		}

		clients = append(clients, c)
	}
}
