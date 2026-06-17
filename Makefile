.PHONY: run build build-linux build-windows build-darwin stop clean

BINARY = voca
VERSION = $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS = -ldflags="-X main.Version=$(VERSION)"

build:
	go build $(LDFLAGS) -o $(BINARY) ./cmd/voca/

build-linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-linux-amd64 ./cmd/voca/

build-windows:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-windows-amd64.exe ./cmd/voca/

build-darwin:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-darwin-amd64 ./cmd/voca/

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
