package main

import (
	"fmt"
	"go_dfs/p2p"
	"log"
)

func main() {

	tr := p2p.NewTCPTransport(p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Decoder:       p2p.NOPDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	})

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("message: %+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
