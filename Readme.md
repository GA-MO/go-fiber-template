# 🚀 Golang Template with Fiber

This is a basic Go project template using the Fiber web framework.

## 📛 Naming Conventions

- [Architecture Guidelines](./Architecture-guidelines.md)
- [Naming Conventions](./Naming-conventions.md)
- [Testing Standards](./Testing-standard.md)

## 🛠️ Prerequisites

- 🦫 Go 1.21 or later

## 🚦 Getting Started

1. 📦 Install dependencies:
```bash
go mod download
```

2. ▶️ Run the server:
```bash
make run
```

3. 🏗️ Build the server:
```bash
make build
```

4. 🚀 Run the server:
```bash
make run-build
```

The server will start on port 9090. You can test it by visiting:
- 🩺 Health check: http://localhost:9090/livez

## 🗂️ Project Structure

- `main.go` - Main application file with server setup
- `go.mod` - Go module file with dependencies
- `Makefile` - Makefile for the project

## 🌐 Available Endpoints

- `GET /livez` - Health check endpoint 
- `GET /readyz` - Ready check endpoint