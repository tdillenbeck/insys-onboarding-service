package responsewrapper

type BadRecord struct {
	ID     string
	Reason string
	Detail string
}

const (
	Duplicate          = "duplicate"
	Invalid            = "invalid"
	MissingInformation = "missing information"
)
