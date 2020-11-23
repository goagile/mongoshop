.DEFAULT = shop
shop:
	cd ./cmd/shop && \
		go build -o ../../shop

.PHONY: run
run:
	./shop -a 127.0.0.1:8081

.PHONY: docs
docs:
	rm -r ./api/docs
	cd ./api && \
		swag init -g ../internal/server/server.go
