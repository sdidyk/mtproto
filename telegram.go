package main

import (
	"fmt"
	"github.com/sdidyk/telegram/mtproto"
	"os"
)

func usage() {
	fmt.Print("Telegram is a simple MTProto tool.\n\nUsage:\n\n")
	fmt.Print("    ./telegram <command> [arguments]\n\n")
	fmt.Print("The commands are:\n\n")
	fmt.Print("    auth     <phonenumber>   authes connection by sms-code\n")
	fmt.Print("    import   <phonenumber>   add phone number in contact list\n")
	fmt.Print("    list                     get contact list\n")
	fmt.Println()
}

func main() {
	var err error

	if len(os.Args) < 2 {
		usage()
		os.Exit(-1)
	}

	commands := map[string]int{"auth": 1, "import": 1, "list": 0}
	valid := false
	for k, v := range commands {
		if os.Args[1] == k {
			if len(os.Args) < v+2 {
				usage()
				os.Exit(-1)
			}
			valid = true
			break
		}
	}

	if !valid {
		usage()
		os.Exit(-1)
	}

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

	switch os.Args[1] {
	case "auth":
		err = m.Auth(os.Args[2])
		if err != nil {
			fmt.Printf("auth: %s\n", err)
			os.Exit(1)
		}
	case "import":
	case "list":
		err = m.GetContacts()
		if err != nil {
			fmt.Printf("list: %s\n", err)
			os.Exit(1)
		}
	}

	// m.Halt()
}
