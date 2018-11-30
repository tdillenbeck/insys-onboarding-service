package uuid

/***************
 * Date: 14/02/14
 * Time: 7:44 PM
 ***************/

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
)

const (
	length = 16

	// 3F used by RFC4122 although 1F works for all
	variantSet = 0x3F

	// rather than using 0xC0 we use 0xE0 to retrieve the variant
	// The result is the same for all other variants
	// 0x80 and 0xA0 are used to identify RFC4122 compliance
	variantGet = 0xE0
)

var (
	// nodeID is the default Namespace node
	nodeId = []byte{
		// 00.192.79.212.48.200
		0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8,
	}
	// The following standard UUIDs are for use with V3 or V5 UUIDs.
	NamespaceDNS  = Namespace{0x6ba7b810, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}
	NamespaceURL  = Namespace{0x6ba7b811, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}
	NamespaceOID  = Namespace{0x6ba7b812, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}
	NamespaceX500 = Namespace{0x6ba7b814, 0x9dad, 0x11d1, 0x80, 0xb4, nodeId, length}
)

// NewV3 will generate a new RFC4122 version 3 UUID
// V3 is based on the MD5 hash of a namespace identifier UUID and
// any type which implements the UniqueName interface for the name.
// For strings and slices cast to a Name type
func NewV3(pNs Namespace, pName UniqueName) UUID {
	var o UUID
	// Set all bits to MD5 hash generated from namespace and name.
	Digest(&o, pNs, pName, md5.New())
	o.setRFC4122Variant()
	o.setVersion(3)
	return o
}

// NewV4 will generate a new RFC4122 version 4 UUID
// A cryptographically secure random UUID.
func NewV4() UUID {
	var o UUID
	// Read random values (or pseudo-randomly) into Array type.
	_, err := rand.Read(o[:length])
	if err != nil {
		panic(err)
	}
	o.setRFC4122Variant()
	o.setVersion(4)
	return o
}

// NewV5 will generate a new RFC4122 version 5 UUID
// Generate a UUID based on the SHA-1 hash of a namespace
// identifier and a name.
func NewV5(pNS Namespacer, pName UniqueName) UUID {
	var o UUID
	Digest(&o, pNS, pName, sha1.New())
	o.setRFC4122Variant()
	o.setVersion(5)
	return o
}

func NewV5FromString(pNS Namespacer, name string) UUID {
	return NewV5(pNS, stringer(name))
}

type stringer string

func (s stringer) String() string {
	return string(s)
}
