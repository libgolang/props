build:
	go build

check: $(GOPATH)/bin/deadcode $(GOPATH)/bin/errcheck $(GOPATH)/bin/golint
	# Run Code Tests
	go vet
	errcheck
	golint
	deadcode

unit-test:
	go test

deploy: build
	

$(GOPATH)/bin/deadcode:
	go get -u github.com/tsenart/deadcode
$(GOPATH)/bin/errcheck:
	go get -u github.com/kisielk/errcheck
$(GOPATH)/bin/golint:
	go get -u golang.org/x/lint/golint

