.PHONY:run
run:
	docker run -p 8080:8081 -it task3

.PHONY:build
build:
	docker build -t task3 .

.PHONY:test
test:
	go test -v

.DEFAULT_GOAL := run