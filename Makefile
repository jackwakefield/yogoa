.PHONY: all dependencies gen test
all: gen test

dependencies:
	gem install watir
	go get -u github.com/xlab/c-for-go
	./scripts/download-yoga.sh

generate: generate-cgo generate-tests

generate-cgo:
	c-for-go --ccincl yoga.yml

generate-tests:
	ruby gentest/gentest.rb
	goimports -w yogoa_test.go

test:
	go test
