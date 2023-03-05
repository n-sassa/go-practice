sqlboiler:
	sqlboiler psql --output ./app/infrastructure/postgresql/models --wipe

make-migrations:
	migrate create -ext sql -dir app/infrastructure/postgresql/migrations -seq ${seq}

make-seed:
	migrate create -ext sql -dir app/infrastructure/postgresql/seed -seq ${seq}

migrate:
	migrate -database="postgres://app_user:password@postgres:5432/develop?sslmode=disable" -path=app/infrastructure/postgresql/migrations/ up
	migrate -database="postgres://app_user:password@postgres:5432/test?sslmode=disable" -path=app/infrastructure/postgresql/migrations/ up

migrate-down:
	migrate -database="postgres://app_user:password@postgres:5432/develop?sslmode=disable" -path=app/infrastructure/postgresql/migrations/ down
	migrate -database="postgres://app_user:password@postgres:5432/test?sslmode=disable" -path=app/infrastructure/postgresql/migrations/ down