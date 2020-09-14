package server

import (
	"reflect"
)

type Server interface {
	Socket() error
	Bind() error
	Serve() error
	Close() error
}

func ListenAndServe(s Server) error {
	defer func() { s.Close() }()

	for _, methodName := range []string{"Socket", "Bind", "Serve"} {
		method := reflect.ValueOf(s).MethodByName(methodName)
		ret := method.Call([]reflect.Value{})
		if e, ok := ret[0].Interface().(error); ok {
			return e
		}
	}

	return nil
}
