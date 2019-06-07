GO = go
RM = rm -rf
CP = cp -t
GOFLAGS = -v
GOBUILD = $(GO) build
BIN_DIR = ./bin
SRC_DIR = ./src
TEST_DIR = ./test

BIN_NAME = huffman-compressor

BIN = $(BIN_DIR)/$(BIN_NAME)

TESTS = $(wildcard $(TEST_DIR)/*)

SHELL = /bin/bash
INSTALL_PREFIX = /usr/local

.PHONY: build clean test install

build:
	$(GOBUILD) -o $(BIN) $(GOFLAGS) $(SRC_DIR)

clean:
	$(RM) $(BIN_DIR)/*
	$(RM) $(TEST_DIR)/*.huff
	$(RM) $(TEST_DIR)/*.ext.*

test:
	@./runtest.sh

install:
	make build
	$(CP) $(INSTALL_PREFIX)/bin $(BIN)

uninstall:
	$(RM) $(INSTALL_PREFIX)/bin/$(BIN_NAME)
