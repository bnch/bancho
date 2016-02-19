package inbound

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetPacket(t *testing.T) {
	r := bytes.NewReader([]byte{0x0C, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x13, 0x76, 0x13, 0x00})
	data, err := GetPacket(r)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Successfully created packet")
	fmt.Printf("ID: %d, Content: % x\n", data.ID, data.Content)
}
