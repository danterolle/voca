.PHONY: run build clean

BINARY = voca

build:
	go build -o $(BINARY) .

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)
