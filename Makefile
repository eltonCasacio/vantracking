createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "postgres://postgres:root@localhost:5432/vantracking?sslmode=disable" -verbose up

migratedown:
	migrate -path=sql/migrations -database "postgres://postgres:root@localhost:5432/vantracking?sslmode=disable" -verbose down

dockerup:
	docker-compose up -d

dockerdown:
	docker-compose down

.PHONY: migrate migrate createmigration dockerup dockerdown