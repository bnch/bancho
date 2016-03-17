package packethandler

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/bnch/bancho/inbound"
	"github.com/bnch/bancho/packethandler/logindata"
	"github.com/bnch/bancho/packets"
	"github.com/bnch/banchoreader/lib"
)

// Handle takes an input and writes data to an output. Not very hard.
func Handle(input []byte, output io.Writer, token string) (string, error) {
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

	var self *Session

	// The user wants to login
	if token == "" {
		sendBackToken = true
		d, err := logindata.Unmarshal(input)
		if err != nil {
			return "", err
		}
		token, deleteAfterwards, err = Login(d)
		if err != nil {
			return token, err
		}
		self = GetSession(token)
	} else if self = GetSession(token); self == nil || self.User.ID == 0 {
		sendBackToken = true
		deleteAfterwards = true
		token = GenerateGUID()
		self = &Session{
			LastRequest: time.Now(),
			stream:      new(bytes.Buffer),
			Mutex:       &sync.Mutex{},
		}
		sessionsMutex.Lock()
		sessions[token] = self
		sessionsMutex.Unlock()
		self.Push(
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
			deleteAfterwards = RawPacketHandler(pack, self)
		}
	}

	// Make up response, putting together all the accumulated packets.
	io.Copy(output, self.stream)

	if deleteAfterwards {
		DeleteCompletely(token)
	}

	if sendBackToken {
		return token, nil
	}
	return "", nil
}
