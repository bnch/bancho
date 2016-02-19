package uleb128

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestMarshal(t *testing.T) {
	printHex(Marshal(1337))
	printHex(Marshal(9001))
	printHex(Marshal(10))
	printHex(Marshal(127))
	printHex(Marshal(128))
	printHex(Marshal(255))
	printHex(Marshal(256))
	printHex(Marshal(624485))
}

func TestUnmarshalReader(t *testing.T) {
	b := bytes.NewReader([]byte("\x93\x03fakestringfakestringfakestring"))
	tot := UnmarshalReader(b)
	d, _ := ioutil.ReadAll(b)
	fmt.Printf("TestUnmarshalReader:\nLength: %d\n%s\n", tot, string(d))
}

func TestUnmarshal(t *testing.T) {
	tUnmarshalSingle(1337, t)
	tUnmarshalSingle(9001, t)
	tUnmarshalSingle(10, t)
	tUnmarshalSingle(127, t)
	tUnmarshalSingle(128, t)
	tUnmarshalSingle(255, t)
	tUnmarshalSingle(256, t)
}
func tUnmarshalSingle(i int, t *testing.T) {
	marshaled := Marshal(i)
	newI, _ := Unmarshal(marshaled)
	if newI != i {
		t.Fatalf("Unmarshaled value (%d) isn't equal to the initial value (%d).", newI, i)
	}
	fmt.Printf("%d ==> success\n", i)
}

func printHex(b []byte) {
	fmt.Printf("% x\n", b)
}
