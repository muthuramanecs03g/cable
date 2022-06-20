package pfcpType

import (
	"encoding/binary"
	"fmt"
	"math"
)

type UPFDataF struct {
	UpfdataValueF float32
}

func (u *UPFDataF) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 8
	data = make([]byte, 4)
	bits := math.Float32bits(u.UpfdataValueF)
	binary.BigEndian.PutUint32(data[idx:], bits)

	return data, nil
}

func (u *UPFDataF) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx + 4 {
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}
	bits := binary.LittleEndian.Uint32(data[idx:])
	u.UpfdataValueF = math.Float32frombits(bits)
	idx = idx + 4

	if length != idx {
		fmt.Printf("%d",length)
		return fmt.Errorf("Inadequate TLV length: %d", length)
	}

	return nil
}
