BIN_DIR := bin/
COV_DIR := cov/
CMD_DIR := ./cmd/
CMD_DIRS := $(patsubst cmd/%/main.go, %, $(wildcard cmd/*/main.go))
CMD_BUILD := $(CMD_DIRS:%=${BIN_DIR}%)
CMD_RUN := $(CMD_DIRS:%=%_run)
TEST_COV := $(CMD_DIRS:%=${COV_DIR}%)
GO_TEST_CMD := go test -bench=. -benchmem -covermode=atomic -race -failfast
GO_PKG_FILES := $(wildcard pkg/*/*.go)
GO_FILES := $(wildcard */*/*.go)
COVERAGE_TXT := coverage.txt

all: run lint test

${BIN_DIR}%: ${CMD_DIR}%/main.go ${GO_PKG_FILES}
	go build -o $@ $<

build: ${CMD_BUILD}

%_run: ${BIN_DIR}%
	time ${BIN_DIR}$* --input ${CMD_DIR}$*/input

run: ${CMD_RUN}

${COV_DIR}%: ${CMD_DIR}%/main_test.go ${BIN_DIR}%
	mkdir -p ${COV_DIR} && ${GO_TEST_CMD} -coverprofile=$@ ${CMD_DIR}$*/...

${COVERAGE_TXT}: ${TEST_COV}
	cat $^ > $@

test: ${COVERAGE_TXT}

lint: ${GO_FILES}
	golangci-lint run -v 2>&1 | tee lint
