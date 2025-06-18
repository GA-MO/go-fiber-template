run:
	go run main.go

test:
	go test ./... -v


build:
	go build -o main main.go

run-build:
	./main

stop:
	@pids=$$(lsof -t -i:9090 2>/dev/null); \
	if [ -n "$$pids" ]; then \
		echo "Killing processes on port 9090: $$pids"; \
		kill -9 $$pids; \
	else \
		echo "No processes found on port 9090"; \
	fi