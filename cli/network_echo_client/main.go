package main

import (
	"fmt"
	"github.com/hachi-n/network_echo/client/echo"
	"os"
)

func main()  {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("arguments error")
		os.Exit(1)
	}

	ipString := "127.0.0.1"

	echo.SendEchoMessage(ipString, args[1])
}
