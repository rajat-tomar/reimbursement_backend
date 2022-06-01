# Reimbursement Backend

This is a reimbursement tool.

## Prerequisites
* [Golang](https://github.com/golang/go) (1.18.2)
* [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
* [Docker](https://docs.docker.com/engine/install/)


## Run Locally

Clone the project

```bash
  git clone git@github.com:gaussb-labs/reimbursement_backend.git
```

Go to the project directory

```bash
  cd reimbursement_backend
```

Install dependencies and add missing module requirements

```bash
  go mod tidy
```
Create config.yml file in project's root directory (refer config.yml.sample)

Make sure your docker is running...

Stop your local postgres service if running
```bash
brew services stop postgresql
```

Start postgres in docker
```bash
  make docker-postgres
```

Create database
```bash
  make create-database
```

Run migration
```bash
  make migrate-up
```

Build project
```bash
  go build
```

Start the server

```bash
  ./reimbursement_backend
```

