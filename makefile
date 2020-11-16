.DEFAULT = server
server: server.go
	go build -o server

.PHONY: run
run:
	./server -a 127.0.0.1:8081
