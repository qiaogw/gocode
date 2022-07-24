build:
	go build -ldflags="-s -w" main.go
	$(if $(shell command -v upx), upx gocode)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o gocode-darwin main.go
	$(if $(shell command -v upx), upx gocode-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o gocode.exe main.go
	$(if $(shell command -v upx), upx gocode.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o gocode-linux main.go
	$(if $(shell command -v upx), upx gocode-linux)

image:
	docker build --rm --platform linux/amd64 -t kevinwan/gocode:$(version) .
	docker tag kevinwan/gocode:$(version) kevinwan/gocode:latest
	docker push kevinwan/gocode:$(version)
	docker push kevinwan/gocode:latest
	docker build --rm --platform linux/arm64 -t kevinwan/gocode:$(version)-arm64 .
	docker tag kevinwan/gocode:$(version)-arm64 kevinwan/gocode:latest-arm64
	docker push kevinwan/gocode:$(version)-arm64
	docker push kevinwan/gocode:latest-arm64
