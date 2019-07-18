package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	. "github.com/logrusorgru/aurora"

	mi "mi/src"
)

func main() {
	for _, f := range os.Args[1:] {
		start := time.Now()
		fmt.Println("Reading begins...\n")
		fmt.Printf("Reading %q after %v %v\n", f, Green(start.Nanosecond()), Green("ns"))
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("Failed to read %q: %v\n", f, err)
			continue
		}
		t, err := mi.ParseTrack(string(d))
		if err != nil {
			fmt.Printf("Failed to parse %q: %v\n", f, err)
			continue
		}
		b, err := t.MarshalBinary()
		if err != nil {
			fmt.Printf("Failed to encode: %v\n", err)
			continue
		}
		err = ioutil.WriteFile(f+".mid", b, 0600)
		if err != nil {
			fmt.Printf("Failed to write %q: %v\n", f+".mid", err)
			continue
		}
	}
}
