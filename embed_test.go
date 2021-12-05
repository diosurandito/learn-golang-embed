package golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed cat.png
var cat []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("cat_new.png", cat, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
