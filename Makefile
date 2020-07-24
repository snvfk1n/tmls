SOURCES := $(find . -name '*.go')

BINARY=build/tmls
LDFLAGS=-ldflags "-X main.BuildTime=`date +%FT%T%z`"
VERSION=$$(git describe --abbrev=0 --tags)

.DEFAULT_GOAL: $(BINARY)

all: clean prebuild test build

.PHONY: prebuild
prebuild: $(SOURCES)
	go get -d -v ./...

.PHONY: build
build: $(SOURCES)
	go build ${LDFLAGS} -o ${BINARY} ${SOURCES}

build-release: clean prebuild test $(SOURCES)
	env GOOS=linux GOARCH=arm go build ${LDFLAGS} -o "${BINARY}-${VERSION}_linux-arm" ${SOURCES}
	env GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o "${BINARY}-${VERSION}_linux-arm64" ${SOURCES}
	env GOOS=linux GOARCH=386 go build ${LDFLAGS} -o "${BINARY}-${VERSION}_linux-386" ${SOURCES}
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o "${BINARY}-${VERSION}_linux-amd64" ${SOURCES}
	env GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o "${BINARY}-${VERSION}_macos-amd64" ${SOURCES}

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: test
test:
	go test
