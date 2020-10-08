.PHONY: build
build: wire
	go build -o ./dist/cli ./cmd

.PHONY: dev
dev:
	CompileDaemon -build="go build -o ./dist/cli ./cmd/cli" -command="./dist/cli.exe -f configs/cli.yml" & \

.PHONY: run
run: wire
	go run ./cmd

.PHONY: wire
wire:
	wire ./...

.PHONY: mock
mock:
	mockery --all
