[![pipeline status](https://gitlab.getweave.com/weave-lab/internal/insys-onboarding/badges/master/pipeline.svg)](https://gitlab.getweave.com/weave-lab/internal/insys-onboarding/commits/master)
[![coverage report](https://gitlab.getweave.com/weave-lab/management/insys-onboarding/badges/master/coverage.svg)](https://gitlab.getweave.com/weave-lab/internal/insys-onboarding/commits/master)

# insys-onboarding

## Installation
```bash
go get weavelab.xyz/insys-onboarding
```

For more information on `weavelab.xyz`, see the projects [readme](https://gitlab.getweave.com/weave-lab/ops/xyz/blob/master/README.md).

## Project Layout

This project's folder structure is based on [Ben Johnson's standard package layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1).
  * `/internal/app` defines the business logic for this service (structs and interfaces)
  * `/internal/config` is used as a default way configure in environment variables.
  * `/internal/grpc`  defines the gRPC handlers.
  * `/internal/mock` defines structs that can be used for testing. This is useful to isolate our unit tests.
  * `/internal/psql` defines code to interact with the database. NOTE: these tests rely on a test database to be setup

## Testing

  This project contains tests that rely on a test database. Here are the setups to setup your local postgres:

  1. psql: `CREATE DATABASE "insys_onboarding_test";`
  2. psql: `CREATE SCHEMA insys_onboarding;` NOTE: make sure to create the schema in the `insys_onboarding_test` database.
  3. `$ make migratetestup`

## Database Migrations
  This service uses the [goose](https://github.com/pressly/goose) library for running migrations. Mainly because it works with schema and doesn't pollute the public namespace.

  Before using goose, the use is responsible for setting up the database:
    1. Install goose `go get -u github.com/pressly/goose/cmd/goose`
    2. psql: `CREATE DATABASE insys_onboarding_local";`
    3. psql: `CREATE SCHEMA insys_onboarding;` NOTE: make sure to create the schema in the `insys_onboarding_local` database.

### Running a migration

  See the Makefile for helper commands to run migrations. Or [RTFM](https://en.wikipedia.org/wiki/RTFM) from the [goose library](https://github.com/pressly/goose).

  Example running migration on local database:
  ```
  make mmigratelocalup
  ```

### Creating a new migration

  ```
  goose -dir dbconfig/migrations/ create MIGRATION_NAME sql
  ```

  Add the SQL for the up migration under the `-- +goose Up` comment. Add the SQL for the down under the `-- +goose Down` comment.

### Seed the database
  The dbconfig/seed.sql contains the seed data for the existing database tables.

  Example running against local database.
  ```
  $ psql postgres://postgres@localhost:5432/insys_onboarding_local?sslmode=disable -f dbconfig/seed.sql
  ```
