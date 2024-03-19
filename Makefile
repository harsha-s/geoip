
.PHONY: mac linux docker run

MAC_OUT_DIR := cmd/mac
LINUX_OUT_DIR := cmd/linux

mac:
	if [ ! -d "$(MAC_OUT_DIR)" ]; then mkdir -p $(MAC_OUT_DIR); fi
	GOOS=darwin GOARCH=amd64 go build -o $(MAC_OUT_DIR)/geoip 

linux:
	if [ ! -d "$(LINUX_OUT_DIR)" ]; then mkdir -p $(LINUX_OUT_DIR); fi
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_OUT_DIR)/geoip 

docker-build:
	docker build -t geoip:latest .		

test:
	go test -v -cover ./...

docker-run:
	docker run -it --env IPSTACK_API_KEY=<API-KEY> geoip:latest -i 8.8.8.8