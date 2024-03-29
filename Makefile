createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/vantracking" -verbose up

migratedown: 
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/vantracking" -verbose down

dockerup:
	docker-compose up -d

dockerdown:
	docker-compose down

.PHONY: migrate migrate createmigration dockerup dockerdown