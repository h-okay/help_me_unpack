package unpack

import (
	"encoding/binary"
	"fmt"
	logger "help_me_unpack/log"
	"math"
)

var UnpackLengthError = "Invalid string to unpack: %s"

type UnpackedBytes struct {
	Integer         int32   `json:"integer,omitempty"`
	UnsignedInteger uint32  `json:"unsigned_integer,omitempty"`
	Short           int16   `json:"short,omitempty"`
	Float           float64 `json:"float,omitempty"`
	Double          float64 `json:"double,omitempty"`
	BigEndianDouble float64 `json:"big___endian___double,omitempty"`
}

func (ub UnpackedBytes) String() string {
	return fmt.Sprintf(
		"Integer:%d, UnsignedInteger:%d, Short:%d, Float:%f, Double:%f, BigEndianDouble:%f",
		ub.Integer,
		ub.UnsignedInteger,
		ub.Short,
		ub.Float,
		ub.Double,
		ub.BigEndianDouble,
	)
}

func Unpack(s string) (UnpackedBytes, error) {
	if len(s) == 0 {
		return UnpackedBytes{}, fmt.Errorf(UnpackLengthError, s)
	}

	decoded, err := Decode([]byte(s))
	if err != nil {
		return UnpackedBytes{}, err
	}

	// Signed/Unsigned Integers and Short
	integer := int32(binary.LittleEndian.Uint32(decoded[0:4]))
	unsignedInteger := binary.LittleEndian.Uint32(decoded[4:8])
	short := int16(binary.LittleEndian.Uint32(decoded[8:12]))

	// Float
	unsignedIntegerTemp := binary.LittleEndian.Uint32(decoded[12:16])
	float := float64(math.Float32frombits(unsignedIntegerTemp))

	// Double
	unsigned64IntegerTemp := binary.LittleEndian.Uint64(decoded[16:24])
	double := math.Float64frombits(unsigned64IntegerTemp)

	// Big Double
	unsigned64IntegerTempBig := binary.BigEndian.Uint64(decoded[24:32])
	bigEndianDouble := math.Float64frombits(unsigned64IntegerTempBig)

	ub := UnpackedBytes{
		Integer:         integer,
		UnsignedInteger: unsignedInteger,
		Short:           short,
		Float:           float,
		Double:          double,
		BigEndianDouble: bigEndianDouble,
	}

	logger.Printf("Unpack result: %s", logger.INFO, ub.String())
	return ub, nil
}
