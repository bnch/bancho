package packethandler

import (
	"bytes"
	"fmt"
	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/packethandler/logindata"
	"io"
	"github.com/bnch/banchoreader/lib"
	"os"
)

// Handle takes an input and writes data to an output. Not very hard.
func Handle(input []byte, output *[]byte, token string) (string, error) {
	sendBackToken := false

	defer func() {
		c := recover()
		if c != nil {
			fmt.Println("ERROR!!!!!!!11!")
			fmt.Println(c)
		}
	}()

	// The user wants to login
	if token == "" {
		sendBackToken = true
		d, err := logindata.Unmarshal(input)
		if err != nil {
			return "", err
		}
		token, err = Login(d, output)
		if err != nil {
			return token, err
		}
	} else {
		inputReader := bytes.NewReader(input)
		for {
			// Find a new packet from input
			pack, err := inbound.GetPacket(inputReader)
			if err != nil && err != io.EOF {
				fmt.Println(err)
			}
			if !pack.Initialised {
				break
			}
			r := banchoreader.New()
			r.Colored = true
			r.DumpPacket(os.Stdout, pack)
		}
	}

	// Make up response, putting together all the accumulated packets.
	for {
		packet := Sessions[token].stream.Pop()
		if packet == nil {
			break
		}
		*output = append(*output, packet.Content...)
	}
	fmt.Printf("% x\n", *output)

	if sendBackToken {
		return token, nil
	}
	return "", nil
}
