package unpack

type UnpackedBytes struct {
	Integer           int32   `json:"integer,omitempty"`
	UnsignedInteger   uint32  `json:"unsigned_integer,omitempty"`
	Short             int16   `json:"short,omitempty"`
	Float             float32 `json:"float,omitempty"`
	Double            float32 `json:"double,omitempty"`
	Big_endian_double float32 `json:"big___endian___double,omitempty"`
}

func Unpack(s string) UnpackedBytes {
	return UnpackedBytes{}
}
