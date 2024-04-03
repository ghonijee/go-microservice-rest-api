DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=user_service


migration-create: 
	migrate create -ext sql -dir database/migrations $(name)

migration-up:
	migrate -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path database/migrations up
migration-rollback:
	migrate -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path database/migrations down
migration-fix:
	migrate -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" -path database/migrations force $(V)