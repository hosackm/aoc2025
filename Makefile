APP_NAME = aoc
CMD_PATH = ./cmd/aoc/main.go
BUILD_DIR = ./bin
BIN = $(BUILD_DIR)/$(APP_NAME)

.PHONY: all build run watch test clean

all: build

build:
	@echo "Building $(APP_NAME)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BIN) $(CMD_PATH)

run:
	@go run $(CMD_PATH)

watch:
	@echo "Watching internal/days for changes..."
	@reflex -r 'internal/days/day[0-9]{2}/.*\.go' -- sh -c '\
		FILE="$${1}"; \
		DAY="$$(basename "$${FILE}" | grep -oE "day[0-9]{2}" | grep -oE "[0-9]{2}")"; \
		if [ -n "$$DAY" ]; then \
			echo "==> Change detected in day $$DAY"; \
			make build; \
			echo "==> Running: $(BIN) run $$DAY"; \
			$(BIN) run $$DAY; \
		fi' sh {}

test:
	@go test ./...

clean:
	@echo "Cleaning build directory"
	@rm -rf $(BUILD_DIR)

