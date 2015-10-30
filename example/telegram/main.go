package main

import (
	"fmt"
	"github.com/sdidyk/mtproto"
	"os"
	"strconv"
)

func usage() {
	fmt.Print("Telegram is a simple MTProto tool.\n\nUsage:\n\n")
	fmt.Print("    ./telegram <command> [arguments]\n\n")
	fmt.Print("The commands are:\n\n")
	fmt.Print("    auth  <phone_number>            auth connection by code\n")
	fmt.Print("    msg	<peer_id> <msgtext>         send message to user\n")
	fmt.Print("    sendmedia <peer_id> <file>      send media file to user\n")
	fmt.Print("    list                            get contact list\n")
	fmt.Print("    dialogs                         get dialogs\n")
	fmt.Println()
}

var commands = map[string]int{
	"auth":      1,
	"msg":       2,
	"sendmedia": 3,
	"list":      0,
	"dialogs":   0,
}

func main() {
	var err error

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	valid := false
	for k, v := range commands {
		if os.Args[1] == k {
			if len(os.Args) < v+2 {
				usage()
				os.Exit(1)
			}
			valid = true
			break
		}
	}

	if !valid {
		usage()
		os.Exit(1)
	}

	m, err := mtproto.NewMTProto(os.Getenv("HOME") + "/.telegram_go")
	if err != nil {
		fmt.Printf("Create failed: %s\n", err)
		os.Exit(2)
	}

	err = m.Connect()
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(2)
	}

	switch os.Args[1] {
	case "auth":
		err = m.Auth(os.Args[2])
	case "msg":
		str_id := os.Args[2]
		if len(str_id) > 0 {
			if str_id[0:1] == "#" {
				chat_id, _ := strconv.Atoi(str_id[1:])
				err = m.SendChatMsg(int32(chat_id), os.Args[3])
			} else if str_id[0:1] == "@" {
				user_id, _ := strconv.Atoi(str_id[1:])
				err = m.SendMsg(int32(user_id), os.Args[3])
			} else {
				user_id, _ := strconv.Atoi(str_id[1:])
				err = m.SendMsg(int32(user_id), os.Args[3])
			}
		}
	case "list":
		err = m.GetContacts()
	case "dialogs":
		err = m.GetChats()
	case "sendmedia":
		chat_id, _ := strconv.Atoi(os.Args[3])
		_, err = m.SendMedia(os.Args[2], int32(chat_id), os.Args[4])
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
