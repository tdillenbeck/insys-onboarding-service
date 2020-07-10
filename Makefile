USERNAME ?= $(shell bash -c 'read -p "Username: " usr; echo $$usr')
PASSWORD ?= $(shell bash -c 'read -s -p "Password: " pwd; echo $$pwd')

default: gotest

tag:
	./dev/tag.sh

gobuild:
	./dev/build

gotest:
	./dev/test-coverage.sh

govet:
	./dev/vet.sh

goerror:
	./dev/errcheck.sh

migratelocalup:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_local?search_path=insys_onboarding&sslmode=disable" up && pg_dump -O -n insys_onboarding -f ./dbconfig/dump.sql --schema-only postgres://localhost:5432/insys_onboarding_local?&sslmode=disable

migratelocaldown:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_local?search_path=insys_onboarding&sslmode=disable" down && pg_dump -n insys_onboarding -f ./dbconfig/dump.sql --schema-only postgres://localhost:5432/insys_onboarding_local?&sslmode=disable

migratetestup:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_test?search_path=insys_onboarding&sslmode=disable" up

migratetestdown:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_test?search_path=insys_onboarding&sslmode=disable" down
	
migratedev:
	@goose -dir ./dbconfig/migrations postgres "postgres://$(USERNAME):$(PASSWORD)@dev-pgsql-service-1a/services?search_path=insys_onboarding&sslmode=disable&role=insys_onboarding" up

migrateprod:
	make migrateprodut
	make migrateprodca1

migrateprodut:
	@goose -dir ./dbconfig/migrations postgres "postgres://$(USERNAME):$(PASSWORD)@pgsql-service-1a/services?search_path=insys_onboarding&sslmode=disable&role=insys_onboarding" up

# NOTE cloud_sql_proxy must be running
migrateprodca1:
	@goose -dir ./dbconfig/migrations postgres "postgres://$(USERNAME):$(PASSWORD)@localhost:5433/services?search_path=insys_onboarding&sslmode=disable&role=insys_onboarding" up

seedlocal:
	psql "postgres://localhost:5432/insys_onboarding_local?sslmode=disable" -f dbconfig/seed.sql

seedtest:
	psql "postgres://localhost:5432/insys_onboarding_test?sslmode=disable" -f dbconfig/seed.sql

seeddev:
	@psql "postgres://$(USERNAME):$(PASSWORD)@dev-pgsql-service-1a/services?sslmode=disable" -f dbconfig/seed.sql

seedprod: 
	make seedprodut
	make seedprodca1

seedprodut:
	@psql "postgres://$(USERNAME):$(PASSWORD)@pgsql-service-1a/services?sslmode=disable" -f dbconfig/seed.sql

seedprodca1:
	@psql "postgres://$(USERNAME):$(PASSWORD)@localhost:5433/services?sslmode=disable" -f dbconfig/seed.sql
