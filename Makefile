clean:
	rm -rf ./build/* || true

build: clean
	go build -o ./build/server ./cmd/app/main.go

air:
	air

test:
	go test ./...

docker-build:
	docker build -t myapp .

docker-run:
	docker-compose up

deploy:
	docker-compose up -d

swagger:
	swag init -g ./cmd/app/main.go -o ./docs

create-migration:
	migrate create -ext sql -dir migrations/ ${NAME}

migrate-up:
	migrate -source file://migrations -database postgres://root:password@localhost:5432/test?sslmode=disable up

migrate-down:
	migrate -source file://migrations -database postgres://root:password@localhost:5432/test?sslmode=disable down 1

migrate-force:
	migrate -source file://migrations -database postgres://root:password@localhost:5432/test?sslmode=disable force ${V}


.PHONY: build test docker-build docker docker-run deploy swagger create-migration migrate-up migrate-down migrate-force