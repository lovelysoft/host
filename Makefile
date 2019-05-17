build_linux:
	@GOOS=linux go build -v -o bin/host_linux .

build_mac:
	@go build -v -o bin/host_mac .

clear:
	@rm -fr bin/*

run:
	@go run main.go