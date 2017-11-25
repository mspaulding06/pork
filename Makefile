all:
	go install ./...

test:
	go test -v ./...

update:
	glide update

.PHONY: all test update