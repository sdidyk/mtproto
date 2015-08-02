package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"github.com/sovietspaceship/mtproto"
)

func main() {
	var err error

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
	
	commands := map[string]int{
		"auth": 1, 
		"msg": 2, 
		"list": 0, 
		"get_full_chat": 1,
		"get_dialogs": 0,
		"help": 0,
		"exit": 0,
	}
	shell := true
	in := bufio.NewReader(os.Stdin)
	for shell {
	    fmt.Print("% ")
	    inputb, _, err := in.ReadLine()
	    input := string(inputb)
		if (input == "") {
			continue
		}
		if err != nil {
		    break
		}
		args := strings.Split(input, " ")
		switch args[0] {
		case "auth":
			err = m.Auth(args[1])
		case "msg":
			user_id, _ := strconv.Atoi(args[1])
			err = m.SendMsg(int32(user_id), args[2])
		case "list":
			err = m.GetContacts()
		case "get_full_chat":
			chat_id, _ := strconv.Atoi(args[1])
			err = m.GetFullChat(int32(chat_id))
		case "get_dialogs":
			err = m.GetDialogs()
		case "help":
			for v, k := range commands {
				fmt.Printf("    %s [%d]\n", v, k)
			}
		case "exit":
			shell = false
		default:
			fmt.Println(args[0], "not found.")
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
