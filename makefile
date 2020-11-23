.PHONY: docs
docs:
	cd cmd/shop && \
	swag init -g main.go -o ../../api/docs

.PHONY: build
build:
	cd ./cmd/shop && \
	go build -o ../../shop

.PHONY: run
run:
	./shop -a 127.0.0.1:8081

clear:
	rm ./shop && \
	rm -r ./api/docs
