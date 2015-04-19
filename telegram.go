package main

import (
	"fmt"
	"github.com/sdidyk/telegram/mtproto"
	"os"
)

func main() {
	var err error

	m, err := mtproto.NewMTProto(os.Getenv("HOME") + "/.telegram_go")
	if err != nil {
		fmt.Printf("Create failed: %s\n", err)
		os.Exit(1)
	}

	err = m.Connect()
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(1)
	}

	err = m.AuthCheckPhone("79197252746")
	if err != nil {
		fmt.Printf("AuthCheckPhone: %s\n", err)
		os.Exit(1)
	}

	m.Halt()
}
