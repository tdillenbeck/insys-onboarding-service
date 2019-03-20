package wiggum

import (
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/werror"
)

//Errors to check
var NotAuthorizedError = werror.Template("not authorized")

// array helper include function
func include(a []string, key string) bool {
	checkMap := make(map[string]struct{})
	for _, k := range a {
		checkMap[k] = struct{}{}
	}
	_, ok := checkMap[key]
	return ok
}

// array helper include function
func includeInt(a []int, key int) bool {
	checkMap := make(map[int]struct{})
	for _, k := range a {
		checkMap[k] = struct{}{}
	}
	_, ok := checkMap[key]
	return ok
}

func includeUUID(a []uuid.UUID, key uuid.UUID) bool {
	checkMap := make(map[uuid.UUID]struct{})
	for _, k := range a {
		checkMap[k] = struct{}{}
	}
	_, ok := checkMap[key]
	return ok
}
