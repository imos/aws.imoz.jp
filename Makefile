install: configure
	appcfg.py update src

configure: src/configure.go

src/configure.go:
	@echo 'src/configure.go is missing.'
	@false
