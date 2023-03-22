package main

import (
	"time"

	"github.com/chaewonkong/go-spinner"
)

func main() {
	go spinner.Spin(100*time.Millisecond, spinner.Arrow)

	time.Sleep(10 * time.Second)
}
