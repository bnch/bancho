package packethandler

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"time"

	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/packethandler/logindata"
	"github.com/bnch/bancho/packets"
	"github.com/bnch/banchoreader/lib"
)

// Handle takes an input and writes data to an output. Not very hard.
func Handle(input []byte, output *[]byte, token string) (string, error) {
	var sendBackToken bool
	var deleteAfterwards bool

	defer func() {
		c := recover()
		if c != nil {
			fmt.Println("ERROR!!!!!!!11!")
			fmt.Println(c)
			fmt.Println(string(debug.Stack()))
		}
	}()

	// The user wants to login
	if token == "" {
		sendBackToken = true
		d, err := logindata.Unmarshal(input)
		if err != nil {
			return "", err
		}
		token, deleteAfterwards, err = Login(d, output)
		if err != nil {
			return token, err
		}
	} else if Sessions[token] == nil || Sessions[token].User.ID == 0 {
		sendBackToken = true
		deleteAfterwards = true
		token = GenerateGUID()
		Sessions[token] = &Session{
			LastRequest: time.Now(),
			stream:      &packetCollection{},
		}
		Sessions[token].Push(
			packets.OrangeNotification("Your session expired. Nothing to worry about - just log in again!"),
			packets.UserID(-1),
		)
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
			deleteAfterwards = RawPacketHandler(pack, Sessions[token])
		}
	}

	// Make up response, putting together all the accumulated packets.
	for {
		packet := Sessions[token].stream.Pop()
		if packet == nil {
			break
		}

		var b bool
		for _, v := range packet.Ignored {
			if v == token {
				b = true
				break
			}
		}
		if b {
			break
		}

		*output = append(*output, packet.Content...)
	}
	fmt.Printf("% x\n", *output)

	if deleteAfterwards {
		delete(Sessions, token)
	}

	if sendBackToken {
		return token, nil
	}
	return "", nil
}
