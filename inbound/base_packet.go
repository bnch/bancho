package inbound

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"git.zxq.co/ripple/go-bancho/packets/uleb128"
	"io"
)

// BasePacket is the most basic type of packet that must then be converted into an actual packet.
type BasePacket struct {
	ID          uint16
	Content     []byte
	Initialised bool
}

// Unmarshal decodes some data from Content and puts it into an interface.
func (b BasePacket) Unmarshal(results ...interface{}) error {
	var err error
	data := bytes.NewReader(b.Content)
	for _, v := range results {
		switch v := v.(type) {

		case *string:
			d, err := readString(data)
			if err != nil {
				return err
			}
			*v = string(d)
		case *[]byte:
			d, err := readString(data)
			if err != nil {
				return err
			}
			*v = d

		case *int8:
			var finVal int8
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal
		case *uint8:
			var finVal uint8
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal

		case *int16:
			var finVal int16
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal
		case *uint16:
			var finVal uint16
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal

		case *int32:
			var finVal int32
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal
		case *uint32:
			var finVal uint32
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal

		case *int64:
			var finVal int64
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal
		case *uint64:
			var finVal uint64
			err = binary.Read(data, binary.LittleEndian, &finVal)
			if err != nil {
				return err
			}
			*v = finVal

		case *[]uint32:
			var arrlen uint16
			err = binary.Read(data, binary.LittleEndian, &arrlen)
			if err != nil {
				return err
			}
			finalArr := make([]uint32, arrlen)
			for i := 0; i < int(arrlen); i++ {
				err = binary.Read(data, binary.LittleEndian, &finalArr[i])
				if err != nil {
					return err
				}
			}
			*v = finalArr

		case *[]int32:
			var arrlen uint16
			err = binary.Read(data, binary.LittleEndian, &arrlen)
			if err != nil {
				return err
			}
			finalArr := make([]int32, arrlen)
			for i := 0; i < int(arrlen); i++ {
				err = binary.Read(data, binary.LittleEndian, &finalArr[i])
				if err != nil {
					return err
				}
			}
			*v = finalArr

		default:
			return fmt.Errorf("packet unmarshal: type not supported (%T)", v)
		}
	}
	return nil
}
func readString(data io.Reader) ([]byte, error) {
	b := make([]byte, 1)
	data.Read(b)
	if b[0] != 11 {
		return []byte{}, errors.New("was expecting string, does not begin with byte 11")
	}
	strlen := uleb128.UnmarshalReader(data)
	b = make([]byte, strlen)
	data.Read(b)
	return b, nil
}
