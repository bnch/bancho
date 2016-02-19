package inbound

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	p := BasePacket{
		Content: []byte("\x00\xa0\x04\x03\x00\x00\x0b\x04Howl"),
	}
	var sho uint16
	var lo uint32
	var st string
	err := p.Unmarshal(&sho, &lo, &st)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sho, lo, st)
}
