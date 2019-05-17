build:
	@echo "=====> Build MacOS Version <====="
	@go build -v -o bin/host_mac .
	@echo "=====> Build Linux Version <====="
	@GOOS=linux go build -v -o bin/host_linux .
	@echo "=====> Completed ! <====="

clear:
	@rm -fr bin/*