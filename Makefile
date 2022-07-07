.PHONY: build

EXIT=exit 0

run: build
	./build/apiserver || $(EXIT)

build:
	go build -o build/ -v ./cmd/apiserver

migrate:
	migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" down
	migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" up

drop:
	migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" drop
	migrate -path migrations -database "postgres://localhost/legko?sslmode=disable" up

.DEFAULT_GOAL := run