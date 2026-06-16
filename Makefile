.PHONY: run build stop clean

BINARY = voca

build:
	go build -o $(BINARY) .

run: build
	@if pgrep -q ollama; then \
		./$(BINARY); \
	else \
		ollama serve >/dev/null 2>&1 & \
		sleep 2; \
		./$(BINARY); \
		pkill ollama 2>/dev/null; \
		echo "Ollama stopped."; \
	fi

stop:
	-pkill -f "$(BINARY)"
	-pkill ollama
	@echo "Stopped."

clean:
	rm -f $(BINARY)
