test: install-deps
	@clear
	cd cli && go test

install-deps:
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega

install: install-deps
	go install
