install: configure
	appcfg.py update src

test: configure
	GOPATH="$$(pwd)" goapp test src/*.go

configure: src/configure.go

src/configure.go:
	@echo 'src/configure.go is missing.'
	@false
