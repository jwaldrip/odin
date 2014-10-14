test:
	@clear
	goop exec sh -c "cd cli && go test"
	go build ./examples/greet-with && greet-with -c red -l hello to world

install-deps:
	goop install

install: install-deps
	goop exec go install
