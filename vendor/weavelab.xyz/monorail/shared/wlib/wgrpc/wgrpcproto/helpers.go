// wgrpcproto is deprecated, please use proto repo shared

package wgrpcproto

import (
	"weavelab.xyz/monorail/shared/wlib/uuid"
	"weavelab.xyz/monorail/shared/wlib/wgrpc/wgrpcproto/wgrpcprotouuid"
)

//UUID converts a *wgrpcprotouuid.UUID into a uuid.UUID, or returns an error if it can't
func UUID(u *wgrpcprotouuid.UUID) (uuid.UUID, error) {
	return u.UUID()
}

//UUIDProto converts a uuid.UUID into a *wgrpcprotouuid.UUID
func UUIDProto(u uuid.UUID) *wgrpcprotouuid.UUID {
	return &wgrpcprotouuid.UUID{
		Bytes: u.Bytes(),
	}
}
