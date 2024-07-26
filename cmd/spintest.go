package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/rmasci/verbose"
)

func main() {
	x := 1
	for i := range x {
		sp := verbose.NewSpinner("Sleep 5:", "stderr", i)
		sp.Speed = 10
		fmt.Printf("%d - %s\n", i, strings.Join(sp.Chars, " "))
		go sp.Start()
		time.Sleep(2 * time.Second)
		sp.Quit <- true
	}
	fmt.Println("Verbose")
	var verb verbose.Verb
	verb.V = true
	go verb.Spin()
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
	verb.Quit <- true
	fmt.Println("done")
}
