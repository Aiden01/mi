package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	. "github.com/logrusorgru/aurora"
	"github.com/eagle453/mi/src"
)

func main() {
	for _, f := range os.Args[1:] {
		start := time.Now()
		fmt.Println("Reading begins...\n")
		d, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Printf("%v Failed to read %q: %v\n", (Bold(Red("Error:"))), f, err)
			continue
		}
		t, err := mi.ParseTrack(string(d))
		if err != nil {
			fmt.Printf("%v Failed to parse %q: %v\n", (Bold(Red("Error:"))), f, err)
			continue
		}
		b, err := t.MarshalBinary()
		if err != nil {
			fmt.Printf("%v Failed to encode: %v\n", (Bold(Red("Error:"))), err)
			continue
		}
		err = ioutil.WriteFile(f+".mid", b, 0600)
		if err != nil {
			fmt.Printf("%v Failed to write %q: %v\n", (Bold(Red("Error:"))), f+".mid", err)
			continue
		}
		fmt.Printf("%v Reading %q after %v %v\n", (Bold(Green("Success:"))), f, Green(start.Nanosecond()), Green("ns"))
	}
}
