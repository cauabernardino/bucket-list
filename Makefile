SHELL=/bin/bash

build:
	docker-compose up -d
down:
	docker-compose down -v
dbup:
	docker-compose up -d postgres
dbdown:
	docker-compose down -v
dbstart:
	docker-compose start postgres
dbstop:
	docker-compose stop postgres
migrateup:
	source .env && migrate -database "$${POSTGRES_URL}" -path db/migration up
migratedown:
	source .env && migrate -database "$${POSTGRES_URL}" -path db/migration down
test:
	go test -v -cover ./...

.PHONY: loadenv dbup dbdown dbstart dbstop migrateup migratedown test build down