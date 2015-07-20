APP_ID=aws-imoz-jp

install: configure
	goapp deploy -application=$(APP_ID) -version=master src/

format:
	goapp fmt src/*.go

test: configure
	GOPATH="$$(pwd)" goapp test src/*.go

configure: src/configure.go

src/configure.go:
	@echo 'src/configure.go is missing.'
	@false
