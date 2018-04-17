package uuid

import (
	"encoding/json"
	"fmt"
	"strings"
)

/****************
 * Date: 1/02/14
 * Time: 10:08 AM
 ***************/

const (
	variantIndex = 8
	versionIndex = 6
)

// A clean UUID type for simpler UUID versions
type UUID [length]byte

func (UUID) Size() int {
	return length
}

func (o UUID) Version() int {
	return int(o[versionIndex]) >> 4
}

func (o *UUID) setVersion(pVersion int) {
	o[versionIndex] &= 0x0F
	o[versionIndex] |= byte(pVersion) << 4
}

func (o *UUID) Variant() byte {
	return variant(o[variantIndex])
}

func (o *UUID) setVariant(pVariant byte) {
	_ = setVariant(&o[variantIndex], pVariant)
}

func (o *UUID) Unmarshal(pData []byte) {
	copy(o[:], pData)
}

func (o UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Format(string(CleanHyphen)))
}

func (o *UUID) UnmarshalJSON(pData []byte) error {

	s := strings.Trim(string(pData), "\"")
	u, err := Parse(s)
	if err != nil {
		return fmt.Errorf("uuid parse: %s", err)
	}

	*o = u
	return nil
}

func (o *UUID) Bytes() []byte {
	return o[:]
}

func (o UUID) String() string {
	return formatter(o, format)
}

func (o UUID) Format(pFormat string) string {
	return formatter(o, pFormat)
}

// Set the three most significant bits (bits 0, 1 and 2) of the
// sequenceHiAndVariant equivalent in the array to ReservedRFC4122.
func (o *UUID) setRFC4122Variant() {
	o[variantIndex] &= 0x3F
	o[variantIndex] |= ReservedRFC4122
}

// Marshals the UUID bytes into a slice
func (o *UUID) MarshalBinary() ([]byte, error) {
	return o.Bytes(), nil
}

// Un-marshals the data bytes into the UUID.
func (o *UUID) UnmarshalBinary(pData []byte) error {
	return UnmarshalBinary(o, pData)
}
