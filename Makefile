.PHONY: build
build:
	go build -o main.so -buildmode=c-shared

.PHONY: run
run: build
	ruby main.rb
