
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

deploydev:
	helm upgrade --kube-context dev-ut  insys-onboarding ./charts/insys-onboarding --reset-values -f ./charts/insys-onboarding/values-dev.yaml --namespace=insys

deployprod:
	make deployprodut
	make deployprodca1

deployprodca1:
	helm upgrade --kube-context gke_weave-canada_northamerica-northeast1_ca1 insys-onboarding ./charts/insys-onboarding --reset-values -f ./charts/insys-onboarding/values-ca.yaml --namespace=insys --install

deployprodut:
	helm upgrade --kube-context prod-ut insys-onboarding ./charts/insys-onboarding --reset-values --namespace=insys

migratelocalup:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_local?search_path=insys_onboarding&sslmode=disable" up && pg_dump -n insys_onboarding -f ./dbconfig/dump.sql --schema-only postgres://localhost:5432/insys_onboarding_local?&sslmode=disable

migratelocaldown:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_local?search_path=insys_onboarding&sslmode=disable" down && pg_dump -n insys_onboarding -f ./dbconfig/dump.sql --schema-only postgres://localhost:5432/insys_onboarding_local?&sslmode=disable

migratetestup:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_test?search_path=insys_onboarding&sslmode=disable" up

migratetestdown:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_test?search_path=insys_onboarding&sslmode=disable" down
	
migratedev:
	goose -dir ./dbconfig/migrations postgres "postgres://username:password@dev-pgsql-service-1a/services?search_path=insys_onboarding&sslmode=disable&role=insys_onboarding" up

migrateprod:
	make migrateprodut
	make migrateprodca1

migrateprodut:
	goose -dir ./dbconfig/migrations postgres "postgres://username:password@pgsql-service-1a/services?search_path=insys_onboarding&sslmode=disable&role=insys_onboarding" up

# NOTE cloud_sql_proxy must be running
migrateprodca1:
	goose -dir ./dbconfig/migrations postgres "postgres://username:password@localhost:5433/services?search_path=insys_onboarding&sslmode=disable&role=insys_onboarding" up

seedlocal:
	psql "postgres://localhost:5432/insys_onboarding_local?sslmode=disable" -f dbconfig/seed.sql

seedtest:
	psql "postgres://localhost:5432/insys_onboarding_test?sslmode=disable" -f dbconfig/seed.sql

seeddev:
	psql "postgres://username:password@dev-pgsql-service-1a/services?sslmode=disable" -f dbconfig/seed.sql

seedprod: 
	make seedprodut
	make seedprodca1

seedprodut:
	psql "postgres://username:password@pgsql-service-1a/services?sslmode=disable" -f dbconfig/seed.sql


seedprodca1:
	psql "postgres://username:password@localhost:5433/services?sslmode=disable" -f dbconfig/seed.sql
