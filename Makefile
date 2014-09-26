test: install-deps
	@cd cli && gopm test

install-deps:
	@cd cli && go get -u github.com/gpmgo/gopm
	@cd cli && gopm install

install: install-deps
	@go install
