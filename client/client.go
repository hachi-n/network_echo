package client

import (
	"fmt"
	"reflect"
)

type Client interface {
	Socket() error
	Connect() error
	Send(string) error
	Recv([]byte) ([]byte, error)
	Close() error
}

func Echo(c Client, message string) error {
	defer func() { c.Close() }()

	for _, methodName := range []string{"Socket", "Connect"} {
		method := reflect.ValueOf(c).MethodByName(methodName)
		ret := method.Call([]reflect.Value{})
		if e, ok := ret[0].Interface().(error); ok {
			return e
		}
	}

	err := c.Send(message)
	if err != nil {
		return err
	}

	buf := make([]byte, len(message))
	recvMessage, err := c.Recv(buf)
	if err != nil {
		return err
	}

	fmt.Println(string(recvMessage))
	return nil
}
