package inbound

import (
	"encoding/binary"
	"io"
)

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
	_, _ = i.Read(make([]byte, 1))

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

	b.Content = make([]byte, contentLength)
	_, err = i.Read(b.Content)
	if i := checkErr(err); i == 2 {
		errF = err
		return
	}

	b.Initialised = true

	return
}

func checkErr(e error) byte {
	if e == nil {
		return 0
	}
	if e == io.ErrUnexpectedEOF {
		return 1
	}
	return 2
}
