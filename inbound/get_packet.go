package inbound

import (
	"encoding/binary"
	"fmt"
	"io"
)

// MaximumContentLength is the maximum length an inbound bancho packet can have (set for sanity). 1024 * 1024 * 10 (10 MB)
const MaximumContentLength = 10485760

// GetPacket returns an bancho packet.
func GetPacket(i io.Reader) (b BasePacket, errF error) {
	err := binary.Read(i, binary.LittleEndian, &b.ID)
	if i := checkErr(err); i > 0 {
		if i == 2 {
			errF = err
		}
		return
	}

	// Read a byte and give no fucks if it returns an error
	i.Read(make([]byte, 1))

	var contentLength uint32
	err = binary.Read(i, binary.LittleEndian, &contentLength)
	if i := checkErr(err); i > 0 {
		// You might think I like copypasting code. I don't. I fucking hate boilerplate code.
		// However, this is life.
		if i == 2 {
			errF = err
		}
		return
	}

	if contentLength > MaximumContentLength {
		errF = fmt.Errorf("are you seriously going to make us believe there's a packet which size is more than %d?! (contentLength: %d)", MaximumContentLength, contentLength)
	}

	b.Content = make([]byte, contentLength)
	read, err := i.Read(b.Content)
	if i := checkErr(err); i == 2 {
		errF = err
		return
	}

	if uint32(read) != contentLength {
		errF = fmt.Errorf("bancho protocol violation: expected to read %d bytes, actually read %d (invalid content length)", contentLength, read)
	}

	b.Initialised = true

	return
}

func checkErr(e error) byte {
	if e == nil {
		return 0
	}
	if e == io.EOF {
		return 1
	}
	return 2
}
