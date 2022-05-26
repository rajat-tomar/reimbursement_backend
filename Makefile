docker-postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust -e PGDATA=/var/lib/postgresql/data/pgdata -d postgres