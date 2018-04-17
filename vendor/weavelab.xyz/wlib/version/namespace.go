package version

import (
	"io/ioutil"
	"strings"
)

const namespaceFilename = `/var/run/secrets/kubernetes.io/serviceaccount/namespace`

// Namespace returns the Kubernetes namespace the application is running
// in. If it isn't running in Kubernetes, a not found error will be
// returned. This should only be used for detecting the environment the application
// is running in.
func Namespace() (string, error) {
	out, err := ioutil.ReadFile(namespaceFilename)
	if err != nil {
		return "", err
	}

	namespace := strings.TrimSpace(string(out))

	return namespace, nil
}
