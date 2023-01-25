all: restart

build:
	docker build .
	docker compose up -d --remove-orphans

restart:
	docker compose down
	docker rmi linkshorterapi-api
	docker compose up -d

clean:
	docker compose down
	docker rmi linkshorterapi-api

test:
	go test -v ./internal/handlers/http
	go test  -v ./internal/service