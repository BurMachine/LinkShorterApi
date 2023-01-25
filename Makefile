all: restart


build:
	docker compose up -d

restart:
	docker compose down
	docker rmi linkshorterapi-api
	docker compose up -d
