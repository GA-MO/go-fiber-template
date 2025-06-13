run:
	go run main.go

test:
	go test ./... -v


build:
	go build -o main main.go

run-build:
	./main