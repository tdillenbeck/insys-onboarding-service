// This package provides RFC4122 UUIDs.
//
// NewV1, NewV3, NewV4, NewV5, for generating versions 1, 3, 4
// and 5 UUIDs as specified in RFC-4122.
//
// New([]byte), unsafe; NewHex(string); and Parse(string) for
// creating UUIDs from existing data.
//
// The original version was from Krzysztof Kowalik <chris@nu7hat.ch>
// Unfortunately, that version was non compliant with RFC4122.
// I forked it but have since heavily redesigned it.
//
// The example code in the specification was also used as reference
// for design.
//
// Copyright (C) 2014 twinj@github.com  2014 MIT style licence
package uuid

/****************
 * Date: 31/01/14
 * Time: 3:35 PM
 ***************/

import (
	"bytes"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"math/big"
	"regexp"
	"strings"
)

const (
	ReservedNCS       byte = 0x00
	ReservedRFC4122   byte = 0x80 // or and A0 if masked with 1F
	ReservedMicrosoft byte = 0xC0
	ReservedFuture    byte = 0xE0
	TakeBack          byte = 0xF0
)

const (

	// Pattern used to parse string representation of the UUID.
	// Current one allows to parse string where only one opening
	// or closing bracket or any of the hyphens are optional.
	// It is only used to extract the main bytes to create a UUID,
	// so these imperfections are of no consequence.
	hexPattern = `^(urn\:uuid\:)?[\{(\[]?([A-Fa-f0-9]{8})-?([A-Fa-f0-9]{4})-?([1-5][A-Fa-f0-9]{3})-?([A-Fa-f0-9]{4})-?([A-Fa-f0-9]{12})[\]\})]?$`
)

var (
	parseUUIDRegex = regexp.MustCompile(hexPattern)
	format         string
)

func init() {
	_ = SwitchFormat(CleanHyphen)
}

// New creates a UUID from a slice of bytes.
// Truncates any bytes past the default length of 16
func New(pData []byte) (UUID, error) {
	if len(pData) < length {
		return UUID{}, fmt.Errorf("Invalid length %d < %d", len(pData), length)
	}
	var o UUID
	o.Unmarshal(pData[:length])
	return o, nil
}

// Creates a UUID from a hex string
// Will panic if hex string is invalid - will panic even with hyphens and brackets
// Expects a clean string use Parse otherwise.
func NewHex(pUuid string) (UUID, error) {
	bytes, err := hex.DecodeString(pUuid)
	if err != nil {
		return UUID{}, err
	}
	return New(bytes)
}

// Parse creates a UUID from a valid string representation.
// Accepts UUID string in following formats:
//		6ba7b8149dad11d180b400c04fd430c8
//		6ba7b814-9dad-11d1-80b4-00c04fd430c8
//		{6ba7b814-9dad-11d1-80b4-00c04fd430c8}
//		urn:uuid:6ba7b814-9dad-11d1-80b4-00c04fd430c8
//		[6ba7b814-9dad-11d1-80b4-00c04fd430c8]
//
func Parse(pUUID string) (UUID, error) {
	md := parseUUIDRegex.FindStringSubmatch(pUUID)
	if md == nil {
		return UUID{}, errors.New("uuid.Parse: invalid string")
	}
	return NewHex(md[2] + md[3] + md[4] + md[5] + md[6])
}

// Digest a namespace UUID and a UniqueName, which then marshals to
// a new UUID
func Digest(o *UUID, pNs Namespacer, pName UniqueName, pHash hash.Hash) {
	// Hash writer never returns an error
	_, _ = pHash.Write(pNs.Bytes())
	_, _ = pHash.Write([]byte(pName.String()))
	o.Unmarshal(pHash.Sum(nil)[:o.Size()])
}

// Function provides a safe way to unmarshal bytes into an
// existing UUID.
// Checks for length.
func UnmarshalBinary(o *UUID, pData []byte) error {
	if len(pData) != o.Size() {
		return fmt.Errorf("uuid.UnmarshalBinary: invalid length")
	}
	o.Unmarshal(pData)
	return nil
}

// **********************************************  UUID Names

// A UUID Name is a simple string which implements UniqueName
// which satisfies the Stringer interface.
type Name string

// Returns the name as a string. Satisfies the Stringer interface.
func (o Name) String() string {
	return string(o)
}

// NewName will create a unique name from several sources
func NewName(salt string, pNames ...UniqueName) UniqueName {
	var s string
	for _, s2 := range pNames {
		s += s2.String()
	}
	return Name(s + salt)
}

// UniqueName is a Stinger interface
// Made for easy passing of IPs, URLs, the several Address types,
// Buffers and any other type which implements Stringer
// string, []byte types and Hash sums will need to be cast to
// the Name type or some other type which implements
// Stringer or UniqueName
type UniqueName interface {

	// Many go types implement this method for use with printing
	// Will convert the current type to its native string format
	String() string
}

// **********************************************  UUID Printing

// A Format is a pattern used by the stringer interface with which to print
// the UUID.
type Format string

const (
	Clean   Format = "%x%x%x%x%x%x"
	Curly   Format = "{%x%x%x%x%x%x}"
	Bracket Format = "(%x%x%x%x%x%x)"

	// This is the default format.
	CleanHyphen Format = "%x-%x-%x-%x%x-%x"

	CurlyHyphen   Format = "{%x-%x-%x-%x%x-%x}"
	BracketHyphen Format = "(%x-%x-%x-%x%x-%x)"
	GoIdFormat    Format = "[%X-%X-%x-%X%X-%x]"
)

// Gets the current default format pattern
func GetFormat() string {
	return format
}

// Switches the default printing format for ALL UUID strings
// A valid format will have 6 groups if the supplied Format does not
func SwitchFormat(pFormat Format) error {
	form := string(pFormat)
	if strings.Count(form, "%") != 6 {
		return fmt.Errorf("uuid.switchFormat: invalid formatting")
	}
	format = form
	return nil
}

// Same as SwitchFormat but will make it uppercase
func SwitchFormatUpperCase(pFormat Format) {
	form := strings.ToUpper(string(pFormat))
	_ = SwitchFormat(Format(form))
}

// Compares whether each UUID is the same
func Equal(p1 UUID, p2 UUID) bool {
	return bytes.Equal(p1.Bytes(), p2.Bytes())
}

// Format a UUID into a human readable string which matches the given Format
// Use this for one time formatting when setting the default using SwitchFormat
// is overkill.
func Formatter(pUUID UUID, pFormat Format) (string, error) {
	form := string(pFormat)
	if strings.Count(form, "%") != 6 {
		return "", fmt.Errorf("uuid.Formatter: invalid formatting")
	}
	return formatter(pUUID, form), nil
}

// **********************************************  UUID Versions

type UUIDVersion int

const (
	NONE UUIDVersion = iota
	RFC4122v1NotSupported
	DunnoYetv2
	RFC4122v3
	RFC4122v4
	RFC4122v5
)

// ***************************************************  Helpers

// Retrieves the variant from the given byte
func variant(pVariant byte) byte {
	switch pVariant & variantGet {
	case ReservedRFC4122, 0xA0:
		return ReservedRFC4122
	case ReservedMicrosoft:
		return ReservedMicrosoft
	case ReservedFuture:
		return ReservedFuture
	}
	return ReservedNCS
}

// not strictly required
func setVariant(pByte *byte, pVariant byte) error {
	switch pVariant {
	case ReservedRFC4122:
		*pByte &= variantSet
	case ReservedFuture, ReservedMicrosoft:
		*pByte &= 0x1F
	case ReservedNCS:
		*pByte &= 0x7F
	default:
		return fmt.Errorf("uuid.setVariant: invalid variant mask")
	}
	*pByte |= pVariant

	return nil
}

// format a UUID into a human readable string
func formatter(pUUID UUID, pFormat string) string {
	b := pUUID.Bytes()
	return fmt.Sprintf(pFormat, b[0:4], b[4:6], b[6:8], b[8:9], b[9:10], b[10:pUUID.Size()])
}

func (u UUID) IsEmpty() bool {
	return u[6]&0xf0 == 0
}

func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		if len(src) == 16 {
			err := UnmarshalBinary(u, src)
			if err != nil {
				return err
			}
		} else {
			tmp, err := Parse(string(src))
			if err != nil {
				return err
			}

			*u = tmp

			return nil
		}
	case string:
		tmp, err := Parse(src)
		if err != nil {
			return err
		}

		*u = tmp

		return nil
	}

	return errors.New("unable to scan uuid")
}

// Cmp compares x and y as integer values and returns:
//
//   -1 if x <  y
//    0 if x == y
//   +1 if x >  y
//
func (x *UUID) Cmp(y UUID) int {
	xInt := big.Int{}
	yInt := big.Int{}

	xInt.SetBytes(x.Bytes())
	yInt.SetBytes(y.Bytes())

	return xInt.Cmp(&yInt)
}
