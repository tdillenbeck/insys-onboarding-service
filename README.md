[![pipeline status](https://gitlab.getweave.com/weave-lab/internal/insys-onboarding/badges/master/pipeline.svg)](https://gitlab.getweave.com/weave-lab/internal/insys-onboarding/commits/master)
[![coverage report](https://gitlab.getweave.com/weave-lab/management/insys-onboarding/badges/master/coverage.svg)](https://gitlab.getweave.com/weave-lab/internal/insys-onboarding/commits/master)

# insys-onboarding

## Installation
```bash
go get weavelab.xyz/insys-onboarding
```

For more information on `weavelab.xyz`, see the projects [readme](https://gitlab.getweave.com/weave-lab/ops/xyz/blob/master/README.md).


## Database
This service uses the [goose](https://github.com/pressly/goose) library for running migrations. Mainly because it works with schema and doesn't pollute the public namespace.

Before using goose, the use is responsible for setting up the database:
  1. CREATE DATABASE "insys-onboarding_dev"
  2. CREATE SCHEMA insys_onboarding;
