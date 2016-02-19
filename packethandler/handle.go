package packethandler

import (
	"bytes"
	"fmt"
	"git.zxq.co/ripple/go-bancho/inbound"
	"git.zxq.co/ripple/go-bancho/packethandler/logindata"
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
		token, err = Login(d)
		if err != nil {
			return token, err
		}
	} else {

		inputReader := bytes.NewReader(input)
		for {
			// Find a new packet from input
			pack, err := inbound.GetPacket(inputReader)
			if err != nil {
				fmt.Println(err)
			}
			if !pack.Initialised {
				break
			}
			fmt.Printf("Inbound packet: %d - %s", pack.ID, pack.Content)
		}
	}

	// Make up response, putting together all the accumulated packets.
	for {
		packet := Tokens[token].Pop()
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
