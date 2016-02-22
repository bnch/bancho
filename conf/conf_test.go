package conf

import (
	"os"
	"reflect"
	"testing"
)

func TestGetConf(t *testing.T) {
	conf, err := GetConf([]byte(SampleConf))
	if err != nil {
		t.Fatal(err)
	}
	if conf.SQLInfo.Name != "bancho" {
		t.Fatal("Was expecting default table name to be bancho, got " + conf.SQLInfo.Name)
	}
}

func TestWriteSampleConfigAndParse(t *testing.T) {
	err := WriteSampleConf()
	if err != nil {
		t.Fatal(err)
	}
	c, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	if c.SQLInfo.Name != "bancho" {
		t.Fatal("Was expecting default table name to be bancho, got " + c.SQLInfo.Name)
	}
	err = os.Remove("bancho.ini")
	if err != nil {
		t.Fatal(err)
	}
}
func TestGetCached(t *testing.T) {
	err := WriteSampleConf()
	if err != nil {
		t.Fatal(err)
	}
	c, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	c2, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c, c2) {
		t.Fatal("expected c == c2, they are different instead")
	}
	if c.SQLInfo.Name != "bancho" {
		t.Fatal("Was expecting default table name to be bancho, got " + c.SQLInfo.Name)
	}
	err = os.Remove("bancho.ini")
	if err != nil {
		t.Fatal(err)
	}
}

// tests that should fail
func TestGetConfInvalid(t *testing.T) {
	// can't say I didn't have fun with this
	_, err := GetConf([]byte("A Y Y L M A O O O O O O OO O O O O OO O ,lDA- 'm') SEMICOLON DROP TABLE users--"))
	if err == nil {
		t.Fatalf("expecting GetConf to fail because ini file is invalid, err == nil instead")
	}
}
func TestGetWithNoFile(t *testing.T) {
	cached = nil
	_, err := Get()
	if err == nil {
		t.Fatalf("expecting Get to fail because no file is present, err == nil instead")
	}
}
