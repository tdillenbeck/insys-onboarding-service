package version

import (
	"io/ioutil"
	"os"
	"strings"
)

const namespaceFilename = `/var/run/secrets/kubernetes.io/serviceaccount/namespace`

// Namespace returns the Kubernetes namespace the application is running
// in. If it isn't running in Kubernetes, but is a test run it will try to the TEST_NAMESPACE env var value if set.
// If TEST_NAMESPACE isn't set and the DEVELOPMENT env var is set then the namespace is set to dev.
// If it isn't running in Kubernetes and it is not a test run and the DEVELOPMENT env var is not set then a not found error will be
// returned. This should only be used for detecting the environment the application
// is running in.
func Namespace() (string, error) {
	var namespace string

	out, err := ioutil.ReadFile(namespaceFilename)
	if err != nil {

		if testNamespace := os.Getenv("TEST_NAMESPACE"); testNamespace != "" {
			return testNamespace, nil
		}

		d := os.Getenv("DEVELOPMENT")
		if strings.EqualFold(d, "true") {
			return "dev", nil
		}

		return "", err
	}

	namespace = strings.TrimSpace(string(out))

	return namespace, nil
}
