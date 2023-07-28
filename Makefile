# Default target: build the application
compile:
	go build -o bin/rover cmd/main.go

# Rule to run the application
run: compile
	./bin/rover input.txt

# Clean up generated files
clean:
	rm -f bin/rover

test:
	go test -v ./...
