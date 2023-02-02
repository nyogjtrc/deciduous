
version := $(shell git describe --tags --dirty --match 'v*' || echo 'dev')
buildtime := $(shell date +%Y/%m/%dT%H:%M:%S)
commit := $(shell git rev-parse --short HEAD)

build_flag := "-w \
	-X github.com/nyogjtrc/go-ver.Version=$(version) \
	-X github.com/nyogjtrc/go-ver.BuildAt=$(buildtime) \
	-X github.com/nyogjtrc/go-ver.Commit=$(commit)"

.PHONY: all

lint:
	golangci-lint run

test:
	go test ./... -count=1

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	rm coverage.out

run-version:
	go run -v -ldflags $(build_flag) cmd/main.go version

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deciduous -a -v -ldflags $(build_flag) ./cmd/main.go

tar:
	tar zcvf deciduous.tar.gz ./deciduous

clean:
	rm deciduous deciduous.tar.gz
