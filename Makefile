test: install-deps
	@clear
	goop exec sh -c "cd cli && go test"

install-deps:
	goop install

install: install-deps
	goop exec go install
