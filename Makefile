BIN_DIR := bin/
COV_DIR := cov/
CMD_DIR := ./cmd/
GO_FILES := $(wildcard */*/*.go)
GO_PKG_FILES := $(wildcard pkg/*/*.go)
GO_TEST_FILES := $(wildcard */*/*_test.go)
CMD_DIRS := $(patsubst cmd/%/main.go, %, $(wildcard cmd/*/main.go))
CMD_BUILD := $(CMD_DIRS:%=${BIN_DIR}%)
CMD_RUN := $(CMD_DIRS:%=%_run)
TEST_COV := $(patsubst %_test.go, cov/%, ${GO_TEST_FILES})
GO_TEST_CMD := go test -bench=. -benchmem -covermode=atomic -race -failfast
COVERAGE_UNIT_TXT := coverage-unit.txt
COVERAGE_FUNCTIONAL_TXT := coverage-functional.txt

all: run test lint

${BIN_DIR}%: ${CMD_DIR}%/main.go ${GO_PKG_FILES}
	go build -o $@ $<

build: ${CMD_BUILD}

%_run: ${BIN_DIR}%
	time ${BIN_DIR}$* --input ${CMD_DIR}$*/input

run: ${CMD_RUN}

${COV_DIR}%: %_test.go %.go ${GO_PKG_FILES}
	mkdir -p $(dir $@) && ${GO_TEST_CMD} -coverprofile=$@  ./$(dir $*)...

${COVERAGE_UNIT_TXT}: ${TEST_COV}
	cat $^ > $@

${COVERAGE_FUNCTIONAL_TXT}: ${GO_PKG_FILES}
	${GO_TEST_CMD} -coverprofile=$@  ./...

test: ${COVERAGE_UNIT_TXT} ${COVERAGE_FUNCTIONAL_TXT}

lint: ${GO_FILES}
	golangci-lint run -v 2>&1 | tee lint
