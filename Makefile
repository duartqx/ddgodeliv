.DEFAULT_GOAL := local

all: build up

local:
	set -a && source ./.env && cd ./src && go run -mod=vendor . && cd ..

build:
	docker compose build backend --no-cache

up:
	trap 'make clean' INT; \
		docker compose up backend || true

clean:
	docker compose down --rmi all
