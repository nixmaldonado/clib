APP_NAME := clib
VERSION := 0.3.0

# Directories
BUILD_DIR := build

# Binary names for different architectures
BINARY_AMD64 := $(BUILD_DIR)/$(APP_NAME)-amd64
BINARY_ARM64 := $(BUILD_DIR)/$(APP_NAME)-arm64
UNIVERSAL_BINARY := $(BUILD_DIR)/$(APP_NAME)

# Compressed file name
TAR_FILE := $(BUILD_DIR)/$(APP_NAME).tar.gz
SHA_FILE := $(BUILD_DIR)/$(APP_NAME).sha256

# Default target: Build and package
all: build package

# Build binaries for both architectures
build: $(BINARY_AMD64) $(BINARY_ARM64) $(UNIVERSAL_BINARY)

$(BINARY_AMD64):
	@echo "Building $(APP_NAME) for macOS AMD64..."
	env GOOS=darwin GOARCH=amd64 go build -o $(BINARY_AMD64)

$(BINARY_ARM64):
	@echo "Building $(APP_NAME) for macOS ARM64..."
	env GOOS=darwin GOARCH=arm64 go build -o $(BINARY_ARM64)

# Create a universal binary using lipo
$(UNIVERSAL_BINARY): $(BINARY_AMD64) $(BINARY_ARM64)
	@echo "Creating universal binary..."
	lipo -create -output $(UNIVERSAL_BINARY) $(BINARY_AMD64) $(BINARY_ARM64)

# Package the universal binary into a tar.gz file
package: $(UNIVERSAL_BINARY)
	@echo "Packaging $(APP_NAME) into $(TAR_FILE)..."
	tar -czvf $(TAR_FILE) -C $(BUILD_DIR) $(APP_NAME)

# Generate the SHA256 checksum
checksum: $(TAR_FILE)
	@echo "Generating SHA256 checksum..."
	shasum -a 256 $(TAR_FILE) > $(SHA_FILE)
	@echo "SHA256 checksum written to $(SHA_FILE)"
	cat $(SHA_FILE)

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Create build directory if it doesn't exist
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

.PHONY: all build clean package