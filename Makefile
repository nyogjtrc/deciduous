
version := $(shell git describe --tags --dirty --match 'v*' || echo 'dev')
buildtime := $(shell date +%Y/%m/%dT%H:%M:%S)
commit := $(shell git rev-parse --short HEAD)

build_flag := "-w \
	-X github.com/nyogjtrc/deciduous/internal/ver.Version=$(version) \
	-X github.com/nyogjtrc/deciduous/internal/ver.BuildTime=$(buildtime) \
	-X github.com/nyogjtrc/deciduous/internal/ver.Commit=$(commit)"

.PHONY: all

lint:
	golangci-lint run

test:
	go test ./... -count=1

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	rm coverage.out

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o deciduous -a -v -ldflags $(build_flag) ./cmd/main.go

run:
	go run -v -ldflags $(build_flag)  main.go

run-service:
	go run -v -ldflags $(build_flag)  main.go service

install:
	go install -a -v -ldflags $(build_flag)

tar:
	tar zcvf deciduous.tar.gz ./deciduous

clean:
	rm deciduous deciduous.tar.gz
