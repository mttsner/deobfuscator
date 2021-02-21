package ironbrew

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/yuin/gopher-lua/parse"
)

//go:embed test.lua
var test string

func TestGetVmdata(t *testing.T) {
	chunk, err := parse.Parse(strings.NewReader(test), "")
	if err != nil {
		t.Fatal(err)
	}
	initVmdata()
	data := vmdata{}
	err = data.GetVmdata(chunk)
	t.Logf("%#v", data)
	if err == nil {
		t.Error(err)
	}
}

func TestInitIronbrew(t *testing.T) {
	err := InitIronbrew()
	if err != nil {
		t.Error(err)
	}
}

/*
func TestMatch(t *testing.T) {
	target, _ := parse.Parse(strings.NewReader(strTarget), "")
	pattern, err := parse.Parse(strings.NewReader(strTarget), "")

	success, _, := beautifier.Match(target, pattern)

	if !success
}*/