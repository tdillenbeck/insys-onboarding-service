package uuid

/****************
 * Date: 31/01/14
 * Time: 3:34 PM
 ***************/

type Namespacer interface {
	Bytes() []byte
}

// Struct is used for RFC4122 Version 1 UUIDs
type Namespace struct {
	timeLow              uint32
	timeMid              uint16
	timeHiAndVersion     uint16
	sequenceHiAndVariant byte
	sequenceLow          byte
	node                 []byte
	size                 int
}

func (o Namespace) Size() int {
	return o.size
}

func (o Namespace) Version() int {
	return int(o.timeHiAndVersion >> 12)
}

func (o Namespace) Variant() byte {
	return variant(o.sequenceHiAndVariant)
}

// Sets the four most significant bits (bits 12 through 15) of the
// timeHiAndVersion field to the 4-bit version number.
func (o *Namespace) setVersion(pVersion int) {
	o.timeHiAndVersion &= 0x0FFF
	o.timeHiAndVersion |= (uint16(pVersion) << 12)
}

func (o *Namespace) setVariant(pVariant byte) {
	_ = setVariant(&o.sequenceHiAndVariant, pVariant)
}

func (o *Namespace) Unmarshal(pData []byte) {
	o.timeLow = uint32(pData[3]) | uint32(pData[2])<<8 | uint32(pData[1])<<16 | uint32(pData[0])<<24
	o.timeMid = uint16(pData[5]) | uint16(pData[4])<<8
	o.timeHiAndVersion = uint16(pData[7]) | uint16(pData[6])<<8
	o.sequenceHiAndVariant = pData[8]
	o.sequenceLow = pData[9]
	o.node = pData[10:o.Size()]
}

func (o Namespace) Bytes() (data []byte) {
	data = []byte{
		byte(o.timeLow >> 24), byte(o.timeLow >> 16), byte(o.timeLow >> 8), byte(o.timeLow),
		byte(o.timeMid >> 8), byte(o.timeMid),
		byte(o.timeHiAndVersion >> 8), byte(o.timeHiAndVersion),
	}
	data = append(data, o.sequenceHiAndVariant)
	data = append(data, o.sequenceLow)
	data = append(data, o.node...)
	return
}

// Set the three most significant bits (bits 0, 1 and 2) of the
// sequenceHiAndVariant to variant mask 0x80.
func (o *Namespace) setRFC4122Variant() {
	o.sequenceHiAndVariant &= variantSet
	o.sequenceHiAndVariant |= ReservedRFC4122
}
