package main

import (
	"fmt"
	"github.com/sdidyk/telegram/mtproto"
)

func main() {
	var err error

	m := new(mtproto.MTProto)

	err = m.Connect("149.154.175.50:443")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = m.Handshake()
	if err != nil {
		fmt.Println(err)
		return
	}

	m.Dump()
}
