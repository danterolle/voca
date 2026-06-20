.PHONY: run build build-linux build-windows build-darwin release checksums stop clean

BINARY = loqi
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS = -ldflags="-X github.com/danterolle/loqi/cmd/loqi/commands.Version=$(VERSION)"
BUILD_DIR = build

build:
	go build $(LDFLAGS) -o $(BINARY) ./cmd/loqi/

build-linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-linux-amd64 ./cmd/loqi/

build-windows:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-windows-amd64.exe ./cmd/loqi/

build-darwin:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-darwin-amd64 ./cmd/loqi/

release: clean
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux-amd64 ./cmd/loqi/
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-amd64 ./cmd/loqi/
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-arm64 ./cmd/loqi/
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-windows-amd64.exe ./cmd/loqi/
	$(MAKE) checksums
	@echo "\nRelease $(VERSION) ready in $(BUILD_DIR)/"

checksums:
	cd $(BUILD_DIR) && shasum -a 256 * > checksums.txt && cat checksums.txt

run: build
	./$(BINARY) $(ARGS)

stop:
	@if command -v taskkill >/dev/null 2>&1; then \
		taskkill /F /IM ollama.exe 2>nul || true; \
	elif command -v pkill >/dev/null 2>&1; then \
		pkill -f "$(BINARY)" 2>/dev/null; \
		pkill ollama 2>/dev/null; \
	else \
		echo "No process killer found (pkill/taskkill)"; \
	fi
	@echo "Stopped."

clean:
	rm -f $(BINARY) $(BINARY)-linux-amd64 $(BINARY)-windows-amd64.exe $(BINARY)-darwin-amd64
	rm -rf $(BUILD_DIR)
