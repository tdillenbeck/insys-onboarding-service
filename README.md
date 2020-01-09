# insys onboarding service

The insys-onboarding-service is for managing the onboarding process for new customers.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

- Install [PostgreSQL](https://www.postgresql.org/) - Relational Database System
- Install [NSQ](https://nsq.io/) - Realtime Distributed Messaging Platform
- Install [helm](https://helm.sh/) for deploys
- Install [goose](https://github.com/pressly/goose) - Database migrations

### Prerequisites

This project's folder structure is based on [Ben Johnson's standard package layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1). The domain logic is defined in the `internal/app/` package instead of in the top level directory.

### Installing

```bash
go get -u weavelab.xyz/insys-onboarding-service
```

#### Running Locally

NSQ is not required for the service to be running locally. This service does publish and consume NSQ messages, so if you are wanting that behavior see the `Running NSQ Locally` step.

##### Setup Local Database
    1. psql: `CREATE DATABASE insys_onboarding_local;`
    2. psql: `CREATE SCHEMA insys_onboarding;` NOTE: create this in the insys_onboarding_local database.
    3. `make migratelocalup`
    4. psal: `CREATE DATABASE insys_onboarding_test;`
    5. psql: `CREATE SCHEMA insys_onboarding;` NOTE: create this in the insys_onboarding_test database.
    6. `make migratetestup`

##### Running NSQ Locally

Follow the quick start guide found [here](https://nsq.io/overview/quick_start.html). NOTE: when starting nsqd, run `nsqd --lookupd-tcp-address=127.0.0.1:4160 -broadcast-address=127.0.0.1`.

##### Running the service Locally

Use the `run` script to run the server locally. This script will setup environment variables to allow the local service to connect with other services.

```bash
./run -d
```

Using the `-d` flag will connect the local service running to services running in the dev kubernetes environment. You can also use the `-p` flag to have the local service connect to production kubernetes environment.

## Running the tests

The postgres tests require that the test database is setup and running. See `Setup Local Database` task on how to do this.

```bash
go test ./...
```

### Break down into end to end tests

Currently there are no end to end tests.

## Creating a database migration

This project uses the [goose](https://github.com/pressly/goose) library to manage migrations. Migrations are created in the `dbconfig/migrations/` folder.

  ```bash
  goose -dir dbconfig/migrations/ create MIGRATION_NAME sql
  ```

Add the SQL for up under the `-- +goose Up` comment. Add the SQL for the down under the `-- +goose Down` comment.

### Run the migrations locally

```bash
make migratelocalup
```

### Run the migrations in dev environment

Add your database credentials to the Makefile for the `migratedev` task (make sure you don't commit them).

```bash
make migratedev
```

### Run the migrations in prod environment

Add your database credentials to the Makefile for the `migrateprod` task (make sure you don't commit them).

```bash
make migrateprod
```

## Deployment

This service uses [helm](https://github.com/helm/charts) to manage our kubernetes deploys. The help charts can be found in the `charts/` folder.

[Kubernetes](https://kubernetes.io/) - System for automating deployment, scaling, and management of containerized applications
[Helm](https://github.com/helm/helm) - Kubernetes Package Manager
[Jenkins](https://jenkins.io/) - Automated build server
[Weave Jenkins Build Server](https://builds.weavelab.xyz/) - UI to view Weave builds.

Use [bart](https://github.com/weave-lab/bart) to get the build tags from the CI/CD server.

```bash
$ bart builds
```

### Deploy to dev environment

If you have added database migrations, make sure to use the `make migratedev` task before deploying.

Update the build tag value in `charts/insys-onboarding/values-dev.yaml` file.

```bash
make deploydev
```

View the kubernetes dashboard for the dev environment at http://dev-dashboard.weave.local/.

### Deploy to prod environment

If you have added database migrations, make sure to use the `make migrateprod` task before deploying.

Update the build tag in `charts/insys-onboarding/values.yaml` file.

```bash
make deployprod
```
View the kubernetes dashboard for the prod environment at http://dashboard-ut.weave.local/.

## Built With

* [Go](https://golang.org/) - Programming Language
* [Vendor](https://github.com/kardianos/govendor) - Dependency Management
