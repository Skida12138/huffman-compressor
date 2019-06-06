GO = go
GOFLAGS = -v
GOBUILD = $(GO) build
BIN_DIR = ./bin
SRC_DIR = ./src
TEST_DIR = ./test

BIN_NAME = huffman-compressor

BIN = $(BIN_DIR)/$(BIN_NAME)

TESTS = $(wildcard $(TEST_DIR)/*)

.PHONY: build clean test

build:
	$(GOBUILD) -o $(BIN) $(GOFLAGS) $(SRC_DIR)

clean:
	rm -rf $(BIN_DIR)/*
	rm -rf $(TEST_DIR)/*.huff
	rm -rf $(TEST_DIR)/*.ext.*

test:
	$(foreach test, $(TESTS), $(BIN) -src $(test) -dst $(test).huff)
	$(foreach test, $(TESTS), $(BIN) -ext -src $(test).huff -dst $(basename $(test)).ext$(suffix $(test)))
	$(foreach test, $(TESTS), if diff $(test) $(basename $(test)).ext$(suffix $(test));\
														then echo "Test case $(basename $(notdir $(test))) passed";\
														else echo "Test case $(basename $(notdir $(test))) faild"; fi;)
