test: install-deps
	@cd cli && $$GOPATH/bin/gopm test

install-deps:
	@cd cli && go get -u github.com/gpmgo/gopm
	@cd cli && $$GOPATH/bin/gopm install

install: install-deps
	@go install
