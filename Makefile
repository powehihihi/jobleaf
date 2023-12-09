BIN_PATH=bin/resume-cli
MAIN_PATH=main.go


.PHONY: build
build:
	go build -o $(BIN_PATH) $(MAIN_PATH)

.PHONY: run
run:
	go run $(MAIN_PATH)

.PHONE: lint
lint:
	gofumpt -l -w . && golangci-lint run -v ./...
