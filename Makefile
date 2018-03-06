test:
	@clear
	@cd cli && go test
	@go build ./examples/greet-with && ./greet-with -c red -l hello to world

install-deps:
	@dep ensure

install:
	@go install
