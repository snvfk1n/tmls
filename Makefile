SOURCEDIR=.
SOURCES := $(find $(SOURCEDIR) -name '*.go')

BINARY=build/tmls
LDFLAGS=-ldflags "-X main.BuildTime=`date +%FT%T%z`"

.DEFAULT_GOAL: $(BINARY)

all: clean test build

.PHONY: build
build: $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} ${SOURCES}

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: test
test:
	go test
