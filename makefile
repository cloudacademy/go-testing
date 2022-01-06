clean:
	rm -rf $(CURDIR)/out/bin

build:
	mkdir -p $(CURDIR)/out/bin
	cd cmd/client && GO111MODULE=on go build -o $(CURDIR)/out/bin/rocket .
	tree

buildall:
	mkdir -p $(CURDIR)/out/bin
	echo "Compiling for every OS and Platform"
	cd cmd/client && GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o $(CURDIR)/out/bin/rocket-darwin-and64 .
	cd cmd/client && GO111MODULE=on GOOS=linux GOARCH=arm64 go build -o $(CURDIR)/out/bin/rocket-linux-arm64 .
	cd cmd/client && GO111MODULE=on GOOS=freebsd GOARCH=386 go build -o $(CURDIR)/out/bin/rocket-freebsd-386 .

run:
	cd cmd/client && go run main.go

test:
	go test -v ./...

testcov:
	go test -v ./... -coverprofile=coverage.out

all: clean testcov buildall