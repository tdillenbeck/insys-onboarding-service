package wapp

var globalLivenessChecks []LivenessCheck

// AddLivenessCheck adds a check to the global list of liveness checks that
// the wapp.Probes handler will add. Calling this function is not thread safe.
// Once wapp.Up is called, this function doesn't do anything.
func AddLivenessCheck(f LivenessCheck) {
	globalLivenessChecks = append(globalLivenessChecks, f)
}

// AddLivenessCheckFunc adds a check to the global list of liveness checks that
// the wapp.Probes handler will add. Calling this function is not thread safe.
// Once wapp.Up is called, this function doesn't do anything.
func AddLivenessCheckFunc(f LivenessCheckFunc) {
	AddLivenessCheck(LivenessCheckFunc(f))
}
