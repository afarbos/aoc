CMD_DIRS :=   $(patsubst cmd/%/main.go, %, $(wildcard cmd/*/main.go))
CMD_BUILD := $(CMD_DIRS:%=%_build)
CMD_RUN := $(CMD_DIRS:%=%_run)
CMD_TEST := $(CMD_DIRS:%=%_test)
GO_TEST_CMD := go test -coverprofile=coverage.txt -covermode=atomic -bench=.
BIN_DIR := ./bin
CMD_DIR := ./cmd

all: ${CMD_RUN} lint ${CMD_TEST}

%_build:
	go build -o ${BIN_DIR}/$* ${CMD_DIR}/$*/main.go

build: ${CMD_BUILD}

%_run: %_build
	time ${BIN_DIR}/$* --input ${CMD_DIR}/$*/input

run: ${CMD_RUN}

%_test:
	${GO_TEST_CMD} ${CMD_DIR}/$*/...

test:
	${GO_TEST_CMD} ./...

lint:
	golangci-lint run -v
