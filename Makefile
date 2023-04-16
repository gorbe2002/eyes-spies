SHELL=/bin/bash
OUT= munch.o
all: compile

compile: clean
	@go build -o $(OUT) munch.go

run: compile
	@./$(OUT)

clean:
ifneq (,$(wildcard ./$(OUT)))
	@rm ./$(OUT)
endif

