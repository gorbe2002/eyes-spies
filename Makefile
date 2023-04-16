SHELL=/bin/bash

all: compile

compile: clean
	@go build -o munch munch.go

run: compile
	@./munch

clean:
ifneq (,$(wildcard ./munch))
	@rm ./munch
endif

