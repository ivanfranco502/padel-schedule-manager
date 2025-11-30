.PHONY: run build test clean help

# Default target
help:
	@echo "Padel Schedule Manager - Available targets:";
	@echo "  make run              - Run with default settings (notify mode)";
	@echo "  make run-interactive  - Run in interactive mode";
	@echo "  make run-reserve      - Run in reserve mode";
	@echo "  make run-freeze       - Run in freeze mode";
	@echo "  make build            - Build the application";
	@echo "  make test             - Run tests";

# Run with default settings
run:
	go run main.go

# Run in interactive mode
run-interactive:
	go run main.go -interactive

# Run with specific modes
run-reserve:
	go run main.go -mode=reserve

run-notify:
	go run main.go -mode=notify

run-freeze:
	go run main.go -mode=freeze

# Run with custom config
run-custom:
	@if [ -z "$(CONFIG)" ]; then \
		echo "Usage: make run-custom CONFIG=$(CONFIG)"; \
		exit 1; \
	fi
	go run main.go -config=$(CONFIG)

# Combine flags (example)
run-full:
	go run main.go -interactive -mode=$(MODE) -config=$(CONFIG)

# Build the application
build:
	go build -o padel-schedule-manager main.go

# Run tests
test:
	go test -v ./...