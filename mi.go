package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"src/text"
	"src/util"
	"src/types"
)

func main() {
	fir _, f := range os.Args[1:] {
		fmt.Printf("reading %q\n", f)
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("failed to read %q: %v\n", f, err)
			continue
		}
		t, err := mi.ParseTrack(string(d))
		if err != nil {
			fmt.Printf("failed to parse %q: %v\n", f, err)
			continue
		}
		err =ioutil.WriteFile(f+".mid", t.MarshalBinary(), 0600)
		if err != nil {
			fmt.Printf("failed to write %q: %v\n", f+".mid", err)
			continue
		}
	}
}