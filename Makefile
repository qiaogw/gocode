build:
	go build -ldflags="-s -w" gocode.go
	$(if $(shell command -v upx), upx gocode)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o gocode-darwin gocode.go
	$(if $(shell command -v upx), upx gocode-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o gocode.exe gocode.go
	$(if $(shell command -v upx), upx gocode.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o gocode-linux gocode.go
	$(if $(shell command -v upx), upx gocode-linux)


