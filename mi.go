package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	mi "Mi/mi/mi/Mi/src"
)

func main() {
	for _, f := range os.Args[1:] {
		start := time.Now()
		fmt.Printf("reading %q after %v ns\n", f, start.Nanosecond())
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
		err = ioutil.WriteFile(f+".mid", t.MarshalBinary(), 0600)
		if err != nil {
			fmt.Printf("failed to write %q: %v\n", f+".mid", err)
			continue
		}
	}
}
