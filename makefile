clean:
	rm -rf $(CURDIR)/out/bin

build:
	mkdir -p $(CURDIR)/out/bin
	cd cmd/client && GO111MODULE=on go build -o $(CURDIR)/out/bin/rocket .
	tree

run:
	cd cmd/client && go run main.go

test:
	go test -v ./...

testcov:
	go test -v ./... -coverprofile=coverage.out

all: clean testcov build