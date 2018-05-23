.PHONY: all dependencies gen test
all: gen test

dependencies:
	go get -u github.com/xlab/c-for-go
	./scripts/download-yoga.sh

gen:
	c-for-go --ccincl yoga.yml

test:
	go test
