package main

import (
	"github.com/hachi-n/network_echo/server/echo"
)

func main() {
	ipString := "127.0.0.1"
	queueLimit := 100

	echo.ListenAndServe(ipString, queueLimit)
}
