test:
	@clear
	cd cli && go test
	go build ./examples/greet-with && ./greet-with -c red -l hello to world
	go build ./examples/greet-with && ./greet-with -h
	go build ./examples/greet-with && ./greet-with -c red -l hello two world

install-deps:
	go get github.com/mattn/goveralls
	go get github.com/axw/gocov
	go get golang.org/x/tools/cmd/cover

install:
	go install
