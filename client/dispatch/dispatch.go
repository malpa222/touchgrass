package dispatch

import (
	"touchgrass/client"
)

var _peerId [20]byte
var _infoHash [20]byte
var _workers []worker

func Init(peerId [20]byte, infoHash [20]byte) {
	_peerId = peerId
	_infoHash = infoHash
}

func GetPiece(pieceChan chan client.Piece) {
	//w := worker{
	//	pieceChan: pieceChan,
	//}
}

// A worker is a wrapper over peer which handles the downloads
type worker struct {
	pieceChan chan client.Piece
}

//func (w *worker) start() {
//	piece := <-w.pieceChan
//}

//func (w *dispatch) shakeHands() error {
//	hs := &handshake.Handshake{
//		PeerID:   w.peerID,
//		InfoHash: w.infoHash,
//	}
//
//	// send the handshake to the peer
//	_, err := conn.Write(hs.Serialize())
//	if err != nil {
//		return err
//	}
//
//	// read the incoming handshake
//	var temp [68]byte
//	_, err = conn.Read(temp[:])
//	if err != nil {
//		return err
//	}
//
//	hs2, err := handshake.Deserialize(temp[:])
//	if hs2 == nil {
//		return errors.New("received no handshake")
//	} else if hs.InfoHash != hs2.InfoHash {
//		return errors.New("the handshake is invalid")
//	}
//
//	return nil
//}

//func (c *Client) Connect(peer *tracker.Peer) error {
//	conn, err := net.Dial("tcp", peer.String())
//	if err != nil {
//		return err
//	}
//
//	if err := c.shakeHands(conn); err != nil {
//		return err
//	}
//
//	// Assign the connection to peer
//	c.conn = conn
//	defer conn.Close()
//
//	return nil
//}
//
//
//func (c *Client) readMessage() error {
//	if c.conn == nil {
//		return errors.New("the the client is not connected to a peer")
//	}
//
//	return nil
//}
