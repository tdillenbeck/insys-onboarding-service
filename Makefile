
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
	helm upgrade --kube-context prod-ut insys-onboarding ./charts/insys-onboarding --reset-values --namespace=insys

migratelocal:
	goose -dir ./dbconfig/migrations postgres "postgres://postgres@localhost:5432/insys_onboarding_local?search_path=insys_onboarding&sslmode=disable" up && pg_dump -n insys_onboarding -f ./dbconfig/dump.sql --schema-only postgres://localhost:5432/insys_onboarding_local?&sslmode=disable
	
migratedev:
	goose -dir ./dbconfig/migrations postgres "postgres://username:password@dev-pgsql-service-1a/services?search_path=insys_onboarding&sslmode=disable" up

migrateprod:
	goose -dir ./dbconfig/migrations postgres "prod connection string" up

