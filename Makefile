CMD_DIRS :=   $(patsubst cmd/%/., %, $(wildcard cmd/*/.))
CMD_BUILD := $(CMD_DIRS:%=%_build)
CMD_RUN := $(CMD_DIRS:%=%_run)
CMD_TEST := $(CMD_DIRS:%=%_test)

all: ${CMD_RUN} ${CMD_TEST}

%_build:
	go build -o $* cmd/$*/main.go

build: ${CMD_BUILD}

%_run: %_build
	time ./$*

run: ${CMD_RUN}

%_test:
	go test -cover -bench=. ./cmd/$*/...

test: ${CMD_TEST}
