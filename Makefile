
.PHONY: all test clean

test:
	go install github.com/jstemmer/go-junit-report/v2@latest
	go test -covermode=count -coverpkg=./... -coverprofile cover.out -v 2>&1 ./... |  go-junit-report -set-exit-code > junit.xml
	go tool cover -html cover.out -o cover.html
