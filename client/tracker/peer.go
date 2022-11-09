package tracker

import (
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func (p *Peer) shakeHands() {

}

func (p *Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
