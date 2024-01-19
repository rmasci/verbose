package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rmasci/verbose"
)

type Fruits struct {
	Fruit []struct {
		Name  string  `json:"name"`
		Color string  `json:"color"`
		Price float64 `json:"price"`
	} `json:"fruit"`
}

func main() {
	// Notice the formatting is the same as if you used the Linux date command
	verb := verbose.New(os.Stdout, "%A %B %Y, %I:%M:%S %P %Z")
	verb.V = true
	verb.PrintDate = false
	verb.Printj(pickFruit())
	verb.Println("Done")
}

func pickFruit() Fruits {
	jtxt := `{"fruit":[{"name":"apple","color":"green","price":1.2},{"name":"pear","color":"yellow","price":0.55},{"name":"grape","color":"purple","price":0.20},{"name":"cherry","color":"red","price":0.25},{"name":"banana","color":"yellow","price":0.5},{"name":"kiwi","color":"green","price":1.25}]}`
	var j Fruits
	err := json.Unmarshal([]byte(jtxt), &j)
	if err != nil {
		fmt.Println("ERROR: Unmarshalling json!", err)
		os.Exit(1)
	}
	return j
}
