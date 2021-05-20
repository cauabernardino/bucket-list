dbup:
	docker-compose up -d postgres
dbdown:
	docker-compose down -v
dbstart:
	docker-compose start postgres
dbstop:
	docker-compose stop postgres
migrateup:
	migrate -database "postgresql://bucket_lister:bucket_password@localhost:5432/bucket-list_db?sslmode=disable" -path db/migration up
migratedown:
	migrate -database "postgresql://bucket_lister:bucket_password@localhost:5432/bucket-list_db?sslmode=disable" -path db/migration down
test:
	go test -v -cover ./...

.PHONY: dbup dbdown dbstart dbstop migrateup migratedown test