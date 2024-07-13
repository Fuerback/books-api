img := books

docker-image:
	docker build -t $(img) .

docker-compose: build
	docker-compose up

build:
	go build

run-tests:
	go test ./... -cover