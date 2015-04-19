package main

import (
	"fmt"
	"github.com/sdidyk/telegram/mtproto"
	"os"
)

func main() {
	authkeyfile := os.Getenv("HOME") + "/.telegram_go"
	m, err := mtproto.NewMTProto("149.154.175.50:443", authkeyfile)

	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(1)
	}

	m.Halt()
}
