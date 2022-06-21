docker-postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust -e PGDATA=/var/lib/postgresql/data/pgdata -d postgres
user-reimbursement:
	psql -h localhost -U postgres password=postgres -c "create user reimbursement with password 'password' createdb replication";
database:
	psql -h localhost -U postgres password=postgres -c "create database reimbursement owner 'reimbursement';"
migrate-up:
	migrate -path db/migration -database "postgresql://reimbursement:password@localhost:5432/reimbursement?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://reimbursement:password@localhost:5432/reimbursement?sslmode=disable" -verbose down
test-database:
	psql -h localhost -U reimbursement password=password -c "CREATE DATABASE reimbursement_test owner 'reimbursement';"
migrate-up-test-database:
	migrate -path db/migration -database "postgresql://reimbursement:password@localhost:5432/reimbursement_test?sslmode=disable" -verbose up
migrate-down-test-database:
	migrate -path db/migration -database "postgresql://reimbursement:password@localhost:5432/reimbursement_test?sslmode=disable" -verbose down