# https://docs.coveralls.io/go
coverall:
	@go test -v -covermode=count -coverprofile=coverage.out ./...
	@goveralls -coverprofile=coverage.out -service=travis-ci -repotoken=$(COVERALL_TOKEN)
	@rm -f coverage.out

test:
	@go test ./... -cover

.PHONY: coverall test