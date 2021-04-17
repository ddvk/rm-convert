VERSION :=$(shell git describe --tags)
LDFLAGS := "-s -w -X main.version=$(VERSION)"
OUT_DIR := dist
BINARY := rm-convert
BUILD = go build -ldflags $(LDFLAGS) -o $(@) $(CMD) 
TARGETS := $(addprefix $(OUT_DIR)/$(BINARY)-, x64 armv6 armv7 arm64 win64 macos)
GOFILES := $(shell find . -iname '*.go' ! -iname "*_test.go")

.PHONY: all clean test 
all: $(TARGETS)

build: $(OUT_DIR)/$(BINARY)-x64

$(OUT_DIR)/$(BINARY)-x64:$(GOFILES)
	GOOS=linux $(BUILD)

$(OUT_DIR)/$(BINARY)-armv6:$(GOFILES)
	GOARCH=arm GOARM=6 $(BUILD)

$(OUT_DIR)/$(BINARY)-armv7:$(GOFILES)
	GOARCH=arm GOARM=7 $(BUILD)

$(OUT_DIR)/$(BINARY)-win64:$(GOFILES)
	GOOS=windows $(BUILD)

$(OUT_DIR)/$(BINARY)-arm64:$(GOFILES)
	GOARCH=arm64 $(BUILD)

$(OUT_DIR)/$(BINARY)-macos:$(GOFILES)
	GOOS=darwin $(BUILD)

clean:
	rm -f $(OUT_DIR)/*

test: 
	go test ./...

