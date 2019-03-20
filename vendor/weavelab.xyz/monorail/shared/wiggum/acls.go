package wiggum

import "weavelab.xyz/monorail/shared/wlib/uuid"

// Value - permission value
func (p Permission) Value() int {
	return int(p)
}

const (
	// ACL User Types
	WeaveACLType    = ACLType("weave")
	PracticeACLType = ACLType("practice")
)

type ACLType string

type ACL map[string][]Permission

func (a ACL) Locations() []uuid.UUID {
	keys := make([]uuid.UUID, 0, len(a))
	for k := range a {
		u, err := uuid.Parse(k)
		if err != nil {
			continue
		}
		keys = append(keys, u)
	}
	return keys
}
