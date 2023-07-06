build:
	go build -o myapp

test:
	go test ./...

docker-build:
	docker build -t myapp .

docker-run:
	docker-compose up

deploy:
	docker-compose up -d
