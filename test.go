package main

import (
	"fmt"

	"github.com/sef-computin/snikt/sniff"
)

func main() {
	dev := "wlan0"
  ch := make(chan string)

	go sniff.StartUtil(dev, ch)

	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)
		}
	}()
  

  for {

  }
}
