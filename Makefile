up:
	docker-compose up --build

migrate:
	docker compose run --rm migrate up

migrate_down:
	docker compose run --rm migrate down 1

migrate_down_all:
	docker compose run --rm migrate down