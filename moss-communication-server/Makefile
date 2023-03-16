dep:
	go mod download
	go mod tidy

### moss-communication-server ###
build-moss-communication-server: dep
	rm -rf bin/
	go build -tags staging -o bin/moss-communication-server cmd/main.go

run-moss-communication-server: build-moss-communication-server
	./bin/moss-communication-server

build-moss-communication-server-windows: dep
	- rm -r .\bin
	- mkdir bin
	go build -tags staging -o bin/moss-communication-server.exe cmd/main.go

run-moss-communication-server-windows: build-moss-communication-server-windows
	./bin/moss-communication-server.exe
