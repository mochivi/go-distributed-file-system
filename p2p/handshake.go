package p2p

import "errors"

// ErrInvalidHandshake is returned if the handshake betweem
// the local and remote node could not be estabilished
var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakeFunc is an abstraction over Handshake functions for all use cases
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
