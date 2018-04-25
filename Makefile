
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

