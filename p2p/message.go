package p2p

import "net"

// Message holds any arbitrary data that is being sent over
// each transport between two nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}
