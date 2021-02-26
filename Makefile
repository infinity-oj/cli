BIN_FILE:=cli

.PHONY: build
build: wire
	go env -w GOOS=windows
	go env -w GOARCH=amd64
	go build -o ./dist/$(BIN_FILE)-windows-amd64.exe ./cmd

	go env -w GOOS=darwin
	go env -w GOARCH=amd64
	go build -o ./dist/$(BIN_FILE)-darwin-amd64 ./cmd

	go env -w GOOS=linux
	go env -w GOARCH=amd64
	go build -o ./dist/$(BIN_FILE)-linux-amd64 ./cmd

.PHONY: prod
prod: wire
	go env -w GOOS=windows
	go env -w GOARCH=amd64
	go build -o ./dist/$(BIN_FILE)-windows-amd64.exe -ldflags "-s -w" ./cmd

	go env -w GOOS=darwin
	go env -w GOARCH=amd64
	go build -o ./dist/$(BIN_FILE)-darwin-amd64 -ldflags "-s -w" ./cmd

	go env -w GOOS=linux
	go env -w GOARCH=amd64
	go build -o ./dist/$(BIN_FILE)-linux-amd64 -ldflags "-s -w" ./cmd

#	scp .\dist\cli-windows-amd64.exe ai:~/proj3/assets/cli
#	scp .\dist\cli-darwin-amd64 ai:~/proj3/assets/cli
#	scp .\dist\cli-linux-amd64 ai:~/proj3/assets/cli
.PHONY: dev
dev:
	CompileDaemon -build="go build -o ./dist/cli ./cmd/cli" -command="./dist/cli.exe -f configs/cli.yml" & \

.PHONY: run
run: wire
	go run ./cmd

.PHONY: wire
wire:
	C:\Users\Wycer\go\bin\wire ./...

.PHONY: mock
mock:
	mockery --all
