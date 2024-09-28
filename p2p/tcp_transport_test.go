package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(TCPTransportOpts{
		ListenAddr:    listenAddr,
		Decoder:       NOPDecoder{},
		HandshakeFunc: NOPHandshakeFunc,
	})
	assert.Equal(t, tr.ListenAddr, listenAddr)

	// Assert ListenAndAccept returns nil (no error)
	assert.Nil(t, tr.ListenAndAccept())
}
