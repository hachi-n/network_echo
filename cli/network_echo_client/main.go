package main

import (
	"fmt"
	"github.com/hachi-n/network_echo/client"
	"github.com/hachi-n/network_echo/client/tcp"
	"github.com/hachi-n/network_echo/client/udp"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var tcpFlag, udpFlag bool
	var ip, message string
	var port int

	app := &cli.App{
		Name:    "network_echo_client",
		Usage:   "network_echo_client [-u|-t]",
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
			&cli.StringFlag{
				Name:        "message, m",
				Destination: &message,
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

			var cl client.Client
			switch {
			case tcpFlag:
				cl = tcp.NewTcpClient(ip, port)
			case udpFlag:
				cl = udp.NewUdpClient(ip, port)
			}

			return client.Echo(cl, message)
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
