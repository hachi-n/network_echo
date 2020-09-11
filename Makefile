go-clean:
	rm -fr ./pkg

go-build:
	go build  -o ./pkg/network_echo_client ./cli/network_echo_client
	go build  -o ./pkg/network_echo_server ./cli/network_echo_server

build:
	$(MAKE) go-clean
	$(MAKE) go-build

