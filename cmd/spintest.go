package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/rmasci/verbose"
)

func main() {
	for i := range 21 {
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
	quit := make(chan bool)
	go verb.Spin(quit)
	time.Sleep(5 * time.Second)
	quit <- true
	fmt.Println("done")
}
