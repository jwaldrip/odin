test: install-deps
	@clear
	@cd cli && go test

ci: install-deps
	@go get code.google.com/p/go.tools/cmd/cover
	@go get github.com/mattn/goveralls
	@cd cli && go test -v -covermode=count -coverprofile=../coverage.out
	@$$HOME/gopath/bin/goveralls/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $$COVERALLS_TOKEN

install-deps:
	@go get github.com/onsi/ginkgo
	@go get github.com/onsi/gomega

install: install-deps
	@go install
