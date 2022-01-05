COVERAGE = coverage.out
export GO111MODULE := on

all: build

.PHONY: deps
deps:
	go mod download
	go get github.com/mattn/goveralls
	go get github.com/haya14busa/goverage
	go get github.com/rubenv/sql-migrate/sql-migrate

.PHONY: build
build:
	go build -v

.