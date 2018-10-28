.PHONY: all dependencies generate test
all: generate test

dependencies:
	go get -u github.com/derekparker/delve/cmd/dlv
	go get -u github.com/xlab/c-for-go
	./scripts/download-yoga.sh
	docker build -t yogoa/watir ./test

generate:
	c-for-go --ccincl -out ./pkg ./yoga.yml
	docker run -i -u $(shell id -u) -v $(shell pwd):/yogoa yogoa/watir ./test/gentest.rb
	goimports -w pkg/yogoa/yoga_test.go

test: 
	dlv test github.com/jackwakefield/yogoa/pkg/yogoa
