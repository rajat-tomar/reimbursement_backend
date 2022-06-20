docker-postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust -e PGDATA=/var/lib/postgresql/data/pgdata -d postgres
create-database:
	psql -h localhost -U postgres password=postgres -c 'CREATE DATABASE reimbursement;'
migrate-up:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/reimbursement?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/reimbursement?sslmode=disable" -verbose down
create-test-database:
	psql -h localhost -U postgres password=postgres -c 'CREATE DATABASE reimbursement_test;'