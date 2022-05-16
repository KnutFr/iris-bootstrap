.DEFAULT_GOAL := compile
.PHONY: build
all: clean check build run
check:
	@echo "Checking project";
	sh scripts/check.sh

build:
	@echo "Building project";
	sh scripts/build.sh

run:
	@echo "Running backend";
	sh scripts/run.sh

clean:
	@echo "Cleaning stuff";
	sh scripts/clean.sh