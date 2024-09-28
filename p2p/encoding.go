package p2p

import (
	"encoding/gob"
	"fmt"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type NOPDecoder struct{}

func (dec NOPDecoder) Decode(r io.Reader, rpc *RPC) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)

	if err != nil {
		return err
	}

	// Handle no bytes read case
	if n == 0 {
		return fmt.Errorf("no bytes read from reader")
	}

	// Set the message payload
	rpc.Payload = buf[:n]

	// Return the err of the Read operation directly
	return nil
}
