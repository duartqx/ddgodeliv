.DEFAULT_GOAL := up

all: build up

build:
	docker compose build backend --no-cache

up:
	trap 'make clean' INT; \
		docker compose up backend || true

clean:
	docker compose down --rmi all
