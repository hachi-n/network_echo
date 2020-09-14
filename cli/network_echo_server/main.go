package main

import (
	"fmt"
	"github.com/hachi-n/network_echo/server"
	"github.com/hachi-n/network_echo/server/tcp"
	"github.com/hachi-n/network_echo/server/udp"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var tcpFlag, udpFlag bool
	var ip string
	var port, queueLimit int

	app := &cli.App{
		Name:    "network_echo_server",
		Usage:   "network_echo_server [-u|-t]",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "tcp, t",
				Value:       false,
				Destination: &tcpFlag,
			},
			&cli.BoolFlag{
				Name:        "udp, u",
				Value:       false,
				Destination: &udpFlag,
			},
			&cli.StringFlag{
				Name:        "ip, i",
				Value:       "127.0.0.1",
				Destination: &ip,
			},
			&cli.IntFlag{
				Name:        "queueLimit, q",
				Value:       100,
				Destination: &queueLimit,
			},
			&cli.IntFlag{
				Name:        "port, p",
				Value:       7,
				Destination: &port,
			},
		},
		Action: func(c *cli.Context) error {
			// XOR
			if tcpFlag == udpFlag {
				return fmt.Errorf("Either must be set. [--tcp | --udp]")
			}

			var s server.Server
			switch {
			case tcpFlag:
				s = tcp.NewTcpServer(ip, port, queueLimit)
			case udpFlag:
				s = udp.NewUdpServer(ip, port, queueLimit)
			}

			return server.EchoServe(s)
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
