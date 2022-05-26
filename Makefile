docker-postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust -e PGDATA=/var/lib/postgresql/data/pgdata -d postgres
create-database:
	psql -h localhost -U postgres password=postgres -c 'CREATE DATABASE reimbursement;'